package controller

import (
	"github.com/chengjoey/chatadmin/pkg/ssevent"
)

var streamManager *ssevent.Manager

func init() {
	streamManager = ssevent.NewManager()
	// go streamManager.TimeProvider()
}

func StreamOpenListener(uid string) chan interface{} {
	return streamManager.OpenListener(uid)
}

func StreamCloseListener(uid string, listener chan interface{}) {
	streamManager.CloseListener(uid, listener)
}

func StreamBroadcastUserInfos(text string) {
	streamManager.SubmitUserInfo(text)
}
