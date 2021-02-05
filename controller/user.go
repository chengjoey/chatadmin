package controller

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"nhooyr.io/websocket/wsjson"

	"github.com/chengjoey/chatadmin/global"
	"github.com/go-redis/redis/v7"

	"github.com/globalsign/mgo/bson"

	"github.com/spf13/viper"
	"nhooyr.io/websocket"
)

const (
	RoleAdmin            string = "admin"
	RoleUser             string = "user"
	userChannelFmt       string = "user:%s:channel"  // 用户发布信息的stream
	adminChannelFmt      string = "admin:%s:channel" // 用户接受信息的stream
	RChatUserHash        string = "users-map"        // 用户昵称和uid的对应map
	RChatUserSet         string = "users-set"        // 用户加入聊天系统时的集合
	RChatAdminSet        string = "admin-set"        // 初始化时加入管理员
	RChatUserAdminHash   string = "users-admin"      // 用户和管理员聊天时的绑定关系
	RChatUserEnterAtHash string = "users-enter-at"   // 用户加入聊天室的时间
	RNoSub               int64  = 0
)

var tokenSecret = []byte("12345678abcdefgh")
var AdminUser []string = []string{"admin", "administrator"}
var rdb *redis.Client
var SystemUser = &User{NickName: "system"}

func init() {
	rdb = redis.NewClient(&redis.Options{Addr: global.RedisAddr})
	rdb.SAdd(RChatAdminSet, AdminUser)
	// defer rdb.Close()
}

type User struct {
	UID        string    `json:"uid"`
	NickName   string    `json:"nickname"`
	EnterAt    time.Time `json:"enter_at"`
	SubChannel string    `json:"-"`
	Token      string    `json:"token"`
	Role       string    `json:"role"`
	To         string    `json:"-"`
	PubChannel string    `json:"-"`
	// cache      []*Message

	conn *websocket.Conn
}

type UserInfo struct {
	Nickname string `json:"nickname"`
	Admin    string `json:"admin"`
	IsChated bool   `json:"is_chated"`
	EnterAt  string `json:"enter_at"`
}

func (u *User) Init(conn *websocket.Conn, client string) {
	u.conn = conn
	u.EnterAt = time.Now()
	if u.Role == RoleAdmin {
		u.To = client
		// userUID := u.GetUid(client)
		u.SubChannel = fmt.Sprintf(userChannelFmt, client)
		u.PubChannel = fmt.Sprintf(userChannelFmt, client)
		// u.PubChannel = fmt.Sprintf(adminChannelFmt, client)
		rdb.HMSet(RChatUserAdminHash, client, u.NickName)
		newMsg := NewAdminToUserMsg(u)
		u.Publish(newMsg)
	} else {
		rdb.SAdd(RChatUserSet, u.NickName)
		// u.SubChannel = fmt.Sprintf(adminChannelFmt, u.NickName)
		u.SubChannel = fmt.Sprintf(userChannelFmt, u.NickName)
		u.PubChannel = fmt.Sprintf(userChannelFmt, u.NickName)
	}
	rdb.HMSet(RChatUserEnterAtHash, u.NickName, time.Now().Format("2006-01-02 15:04:05"))
}

func (u *User) RecieveMessage(ctx context.Context) error {
	var (
		receiveMsg map[string]string
		err        error
	)
	for {
		err = wsjson.Read(ctx, u.conn, &receiveMsg)
		if err != nil {
			var closeErr websocket.CloseError
			if errors.As(err, &closeErr) {
				return nil
			} else if errors.Is(err, io.EOF) {
				return nil
			}

			return err
		}

		sendMsg := NewMessage(u, receiveMsg["content"], receiveMsg["send_time"])
		data := map[string]interface{}{"data": sendMsg.Hxe()}
		args := &redis.XAddArgs{Stream: u.PubChannel, ID: "*", Values: data}
		msgID, err := rdb.XAdd(args).Result()
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println("msg id : ", msgID)

		// res, err := rdb.Publish(u.PubChannel, sendMsg.Hxe()).Result()
		// if err != nil {
		// 	return err
		// }
		// if res == RNoSub {
		// 	u.cache = append(u.cache, sendMsg)
		// }
	}
}

func (u *User) SendMessage(ctx context.Context) {
	// var data Message
	// for rMsg := range u.SubChannel {
	// 	json.Unmarshal([]byte(rMsg.Payload), &data)
	// 	err := wsjson.Write(ctx, u.conn, data)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	args := &redis.XReadArgs{Streams: []string{u.SubChannel, "$"}, Block: time.Duration(0), Count: 1}
	for {
		stm := rdb.XRead(args).Val()
		for _, stream := range stm {
			for _, msg := range stream.Messages {
				var sendMsg Message
				data := msg.Values["data"]
				_ = json.Unmarshal([]byte(data.(string)), &sendMsg)
				if sendMsg.User.NickName != u.NickName {
					err := wsjson.Write(ctx, u.conn, &sendMsg)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	}
}

func (u *User) Publish(msg *Message) {
	data := map[string]interface{}{"data": msg.Hxe()}
	args := &redis.XAddArgs{Stream: u.PubChannel, ID: "*", Values: data}
	_, err := rdb.XAdd(args).Result()
	if err != nil {
		log.Println(err)
	}
}

func (u *User) GetAllHistoryMsg(client string) ([]Message, error) {
	var target string
	var historyMsgs []Message
	if u.Role == RoleUser {
		target = fmt.Sprintf(userChannelFmt, u.NickName)
	} else {
		target = fmt.Sprintf(userChannelFmt, client)
	}
	rdbMsgs, err := rdb.XRange(target, "-", "+").Result()
	if err != nil {
		return historyMsgs, err
	}
	for _, data := range rdbMsgs {
		var msg Message
		strData := data.Values["data"]
		_ = json.Unmarshal([]byte(strData.(string)), &msg)
		if msg.User.NickName != "system" && msg.Type == MsgTypeNormal {
			historyMsgs = append(historyMsgs, msg)
		}
	}
	return historyMsgs, nil
}

func (u *User) GetUid(nickname string) string {
	return rdb.HGet(RChatUserHash, nickname).Val()
	// return rdb.Get(nickname).Val()
}

func IsAdministrator(nickname string) bool {
	return rdb.SIsMember(RChatAdminSet, nickname).Val()
}

func (u *User) Leave() {
	// rdb.Del(u.NickName)
	// rdb.HMSet(RChatUserIdHash, u.NickName, false)
	if u.Role == RoleAdmin {
		client := strings.Split(u.SubChannel, ":")[1]
		rdb.HDel(RChatUserAdminHash, client)
	} else {
		leaveMsg := NewLeaveMessage(u)
		u.Publish(leaveMsg)
		rdb.SRem(RChatUserSet, u.NickName)
		rdb.HDel(RChatUserAdminHash, u.NickName)
	}
}

func (u *User) CanEnterRoot(client string) error {
	if u.Role == RoleAdmin {
		if !rdb.SIsMember(RChatUserSet, client).Val() {
			return fmt.Errorf("用户： %s不在线", client)
		}
		if rdb.HExists(RChatUserAdminHash, client).Val() {
			return fmt.Errorf("用户: %s已有管理员匹配", client)
		}
	} else {
		if rdb.SIsMember(RChatUserSet, u.NickName).Val() {
			return fmt.Errorf("用户: %s已在聊天室", u.NickName)
		}
	}
	return nil
}

func NewUser(nickname, role string) User {
	user := User{
		NickName: nickname,
		EnterAt:  time.Now(),
		Role:     role,
	}

	// if role == RoleUser {
	// 	uid := user.GetUid(user.NickName)
	// 	if uid == "" {
	// 		user.UID = bson.NewObjectId().Hex()
	// 		rdb.HMSet(RChatUserHash, user.NickName, user.UID)
	// 	} else {
	// 		user.UID = uid
	// 	}
	// } else {
	// 	user.UID = bson.NewObjectId().Hex()
	// }
	uid := user.GetUid(user.NickName)
	if uid == "" {
		user.UID = bson.NewObjectId().Hex()
		rdb.HMSet(RChatUserHash, user.NickName, user.UID)
	} else {
		user.UID = uid
	}

	user.Token = genAesToken(user.NickName, user.UID, user.Role)

	// if IsAdministrator(user.NickName) {
	// 	user.Role = RoleAdmin
	// 	uid := user.GetUid(user.To)
	// 	user.SubChannel = rdb.Subscribe(fmt.Sprintf(userChannelFmt, uid)).Channel()
	// 	user.PubChannel = fmt.Sprintf(adminChannelFmt, uid)
	// } else {
	// 	rdb.SAdd(RChatUserSet, user.NickName)
	// 	rdb.Set(user.NickName, user.UID, 0)
	// 	user.SubChannel = rdb.Subscribe(fmt.Sprintf(adminChannelFmt, user.UID)).Channel()
	// 	user.PubChannel = fmt.Sprintf(userChannelFmt, user.UID)
	// 	// user.SubChannel = rdb.Subscribe(fmt.Sprintf(userChannelFmt, user.NickName))
	// }
	return user
}

func genToken(uid, nickname string) string {
	secret := viper.GetString("token-secret")
	message := fmt.Sprintf("%s%s%s", nickname, secret, uid)

	messageMAC := macSha256([]byte(message), []byte(secret))

	return fmt.Sprintf("%suid%s", base64.StdEncoding.EncodeToString(messageMAC), uid)
}

func parseTokenAndValidate(token, nickname string) (string, error) {
	pos := strings.LastIndex(token, "uid")
	messageMAC, err := base64.StdEncoding.DecodeString(token[:pos])
	if err != nil {
		return "", err
	}
	uid := token[pos+3:]

	secret := viper.GetString("token-secret")
	message := fmt.Sprintf("%s%s%s", nickname, secret, uid)

	ok := validateMAC([]byte(message), messageMAC, []byte(secret))
	if ok {
		return uid, nil
	}

	return "", errors.New("token is illegal")
}

func macSha256(message, secret []byte) []byte {
	mac := hmac.New(sha256.New, secret)
	mac.Write(message)
	return mac.Sum(nil)
}

func validateMAC(message, messageMAC, secret []byte) bool {
	mac := hmac.New(sha256.New, secret)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func genAesToken(nickname, uid, role string) string {
	plaintext := []byte(fmt.Sprintf("%s@%s@%s", nickname, uid, role))
	block, _ := aes.NewCipher(tokenSecret)
	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, tokenSecret[:blockSize])
	crypted := make([]byte, len(plaintext))
	blockMode.CryptBlocks(crypted, plaintext)
	return base64.StdEncoding.EncodeToString(crypted)
}

func DescryptTokenToUser(token string) (User, error) {
	var u User
	crypted, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return u, fmt.Errorf("token解析错误")
	}
	block, _ := aes.NewCipher(tokenSecret)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, tokenSecret[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	datas := strings.Split(string(origData), "@")
	if len(datas) != 3 {
		return u, fmt.Errorf("token解析错误")
	}
	u.NickName = datas[0]
	u.UID = datas[1]
	u.Role = datas[2]
	return u, nil
}
