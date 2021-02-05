package controller

import (
	"context"
	"encoding/json"
	"fmt"

	"nhooyr.io/websocket/wsjson"
)

type broadcaster struct {
	enteringChannel chan *User
	leavingChannel  chan *User
}

var Broadcaster = &broadcaster{
	enteringChannel: make(chan *User),
	leavingChannel:  make(chan *User),
}

func (b *broadcaster) Start() {
	for {
		select {
		case user := <-b.enteringChannel:
			fmt.Println("new user: ", user)
			b.StreamUserInfos()
		case user := <-b.leavingChannel:
			user.Leave()
			b.StreamUserInfos()
		}
	}
}

func (b *broadcaster) SendWelcomeMsg(user *User, ctx context.Context) {
	if user.Role == RoleUser {
		msg := NewWelcomeMessage()
		if user.conn != nil {
			_ = wsjson.Write(ctx, user.conn, &msg)
		}
		// data := map[string]interface{}{"data": msg.Hxe()}
		// args := &redis.XAddArgs{Stream: user.PubChannel, ID: "*", Values: data}
		// msgID, err := rdb.XAdd(args).Result()
		// if err != nil {
		// 	log.Println(err)
		// 	return
		// }
		// log.Println("msg id : ", msgID)
	}
}

func (b *broadcaster) GetUserList() []UserInfo {
	users := []UserInfo{}
	userSet := rdb.SMembers(RChatUserSet).Val()
	for _, nickanme := range userSet {
		var u UserInfo
		u.Nickname = nickanme
		admin := rdb.HGet(RChatUserAdminHash, nickanme).Val()
		u.Admin = admin
		u.EnterAt = rdb.HGet(RChatUserEnterAtHash, u.Nickname).Val()
		if u.Admin != "" {
			u.IsChated = true
		}
		users = append(users, u)
	}
	return users
}

func (b *broadcaster) StreamUserInfos() {
	infos := b.GetUserList()
	infoStr, _ := json.Marshal(infos)
	StreamBroadcastUserInfos(string(infoStr))
}

func (b *broadcaster) UserEntering(u *User) {
	b.enteringChannel <- u
}

func (b *broadcaster) UserLeaving(u *User) {
	b.leavingChannel <- u
}
