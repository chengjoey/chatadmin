package controller

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cast"
)

const (
	MsgTypeNormal    = iota // 普通 用户消息
	MsgTypeWelcome          // 当前用户欢迎消息
	MsgTypeUserEnter        // 用户进入
	MsgTypeUserLeave        // 用户退出
	MsgTypeError            // 错误消息
)

type Message struct {
	User           *User     `json:"user"`
	Type           int       `json:"type"`
	Content        string    `json:"content"`
	MsgTime        time.Time `json:"msg_time"`
	ClientSendTime time.Time `json:"client_send_time"`
	Url            string    `json:"url"`
}

func NewMessage(user *User, content, clientTime string) *Message {
	message := &Message{
		User:    user,
		Type:    MsgTypeNormal,
		Content: content,
		MsgTime: time.Now(),
	}
	if clientTime != "" {
		message.ClientSendTime = time.Unix(0, cast.ToInt64(clientTime))
	}
	return message
}

func (m *Message) Hxe() string {
	bs, _ := json.Marshal(m)
	return string(bs)
}

func NewWelcomeMessage() *Message {
	message := &Message{
		User:    SystemUser,
		Type:    MsgTypeWelcome,
		Content: "欢迎进入客服系统, 将尽快为您安排客服人员",
		MsgTime: time.Now(),
	}
	return message
}

func NewLeaveMessage(user *User) *Message {
	message := &Message{
		User:    SystemUser,
		Type:    MsgTypeUserLeave,
		Content: fmt.Sprintf("客户: %s已离开当前聊天", user.NickName),
		MsgTime: time.Now(),
	}
	return message
}

func NewAdminToUserMsg(user *User) *Message {
	message := &Message{
		User:    user,
		Type:    MsgTypeUserEnter,
		Content: fmt.Sprintf("您好, 我是客服人员: %s，很高兴为您服务", user.NickName),
		MsgTime: time.Now(),
	}
	return message
}
