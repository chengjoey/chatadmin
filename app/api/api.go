package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/chengjoey/chatadmin/pkg/ssevent"

	"github.com/chengjoey/chatadmin/app/def"
	"github.com/chengjoey/chatadmin/global"
	"nhooyr.io/websocket"

	"github.com/chengjoey/chatadmin/controller"
	"github.com/chengjoey/chatadmin/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.DefaultQuery("token", "")
		if token == "" {
			token = c.Request.Header.Get("Authorization")
		}
		if token == "" {
			c.Error(response.WrapError(nil, "token出错"))
			return
		}
		user, err := controller.DescryptTokenToUser(token)
		if err != nil {
			c.Error(response.WrapError(err, "身份出错"))
			return
		}
		c.Set("user", user)
	}
}

func getUser(c *gin.Context) (*controller.User, error) {
	data, _ := c.Get("user")
	user, ok := data.(controller.User)
	if !ok {
		return &user, fmt.Errorf("获取用户失败")
	}
	return &user, nil
}

func HomeHandle(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Login(c *gin.Context) {
	var li def.UserLogin
	if err := c.ShouldBindJSON(&li); err != nil {
		c.Error(response.WrapError(err, "请求参数格式错误"))
		return
	}
	if l := len(li.Nickname); l < 2 || l > 20 {
		c.Error(response.WrapError(nil, "非法昵称: %s, 昵称长度:2-20", li.Nickname))
		return
	}
	role := controller.RoleUser
	isAdmin := controller.IsAdministrator(li.Nickname)
	if isAdmin {
		role = controller.RoleAdmin
	}
	user := controller.NewUser(li.Nickname, role)
	c.Header("token", user.Token)
	user.Token = ""
	c.JSON(http.StatusOK, gin.H{"value": user, "msg": "登录成功"})
}

func WebSocketHandle(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		c.Error(response.WrapError(nil, "获取用户信息失败"))
		return
	}
	client := c.Query("client")
	if err := user.CanEnterRoot(client); err != nil {
		c.Error(response.WrapError(err, "加入聊天失败"))
		return
	}
	options := websocket.AcceptOptions{InsecureSkipVerify: true}
	conn, err := websocket.Accept(c.Writer, c.Request, &options)
	if err != nil {
		c.Error(err)
		return
	}
	defer conn.Close(websocket.StatusInternalError, "内部出错")

	user.Init(conn, client)
	// token := c.Query("token")
	// nickname := c.Query("nickname")
	// to := c.Query("to")
	// if l := len(nickname); l < 2 || l > 20 {
	// 	wsjson.Write(c.Request.Context(), conn, response.AppError{Msg: "非法名称，名称长度：2-20"})
	// 	conn.Close(websocket.StatusUnsupportedData, "nickname illegal!")
	// 	return
	// }
	// if ok, _ := controller.Broadcaster.IsInRoom(nickname); ok {
	// 	wsjson.Write(c.Request.Context(), conn, response.AppError{Msg: "该名称已经存在"})
	// 	conn.Close(websocket.StatusUnsupportedData, "nickname exists")
	// 	return
	// }

	// user := controller.NewUser(conn, token, nickname, c.Request.RemoteAddr, to)
	// user.Token = ""

	controller.Broadcaster.UserEntering(user)

	go user.SendMessage(c.Request.Context())

	controller.Broadcaster.SendWelcomeMsg(user, c.Request.Context())

	err = user.RecieveMessage(c.Request.Context())

	controller.Broadcaster.UserLeaving(user)

	if err == nil {
		conn.Close(websocket.StatusNormalClosure, "")
	} else {
		conn.Close(websocket.StatusInternalError, "Read from client error")
	}
}

func StreamSSEvent(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		c.Error(response.WrapError(nil, "获取用户信息失败"))
		return
	}
	if user.Role != global.RoleAdmin {
		c.Error(response.WrapError(nil, "权限不足"))
		return
	}
	listener := controller.StreamOpenListener(user.UID)
	defer controller.StreamCloseListener(user.UID, listener)

	clientGone := c.Writer.CloseNotify()
	c.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			fmt.Println("用户离开")
			return false
		case message := <-listener:
			sseMsg, ok := message.(*ssevent.Message)
			if !ok {
				return false
			}
			c.SSEvent("message", sseMsg.Text)
			return true
		}
	})
}

func AdminGetUserList(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		c.Error(response.WrapError(nil, "获取用户信息失败"))
		return
	}
	if user.Role != global.RoleAdmin {
		c.Error(response.WrapError(nil, "权限不足"))
		return
	}
	users := controller.Broadcaster.GetUserList()
	c.JSON(http.StatusOK, gin.H{"msg": "获取用户列表成功", "value": users})
}

func GetHistoryMsg(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		c.Error(response.WrapError(nil, "获取用户信息失败"))
		return
	}
	client := c.Query("client")
	msgs, err := user.GetAllHistoryMsg(client)
	if err != nil {
		c.Error(response.WrapError(err, "获取历史消息列表失败"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "获取历史消息列表成功", "value": msgs})
}
