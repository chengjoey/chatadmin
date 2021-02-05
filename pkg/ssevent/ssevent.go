package ssevent

import (
	"time"

	"github.com/dustin/go-broadcast"
)

const (
	Msg_Flag_User_info = 1 //用于想管理员用户实时发送用户在线情况
)

type Message struct {
	Flag uint
	Uid  string
	Text string
}

type Listener struct {
	Channel chan interface{}
	Uid     string
}

type uidRequest struct {
	response chan<- map[string]bool
}

type Manager struct {
	bcChannel broadcast.Broadcaster
	open      chan *Listener
	close     chan *Listener
	delete    chan struct{}
	messages  chan *Message
	uids      map[string]bool
	userChan  chan uidRequest
}

func NewManager() *Manager {
	manager := &Manager{
		bcChannel: broadcast.NewBroadcaster(10),
		open:      make(chan *Listener, 100),
		close:     make(chan *Listener, 100),
		delete:    make(chan struct{}, 100),
		messages:  make(chan *Message, 100),

		uids:     map[string]bool{},
		userChan: make(chan uidRequest, 10),
	}

	go manager.run()
	return manager
}

func (m *Manager) run() {
	for {
		select {
		case listener := <-m.open:
			m.register(listener)
		case listener := <-m.close:
			m.deregister(listener)
		case <-m.delete:
			m.deleteBroadcast()
		case message := <-m.messages:
			m.bcChannel.Submit(message)
		case req := <-m.userChan:
			req.response <- m.uids
		}
	}
}

func (m *Manager) register(listener *Listener) {
	m.uids[listener.Uid] = true
	m.bcChannel.Register(listener.Channel)
}

func (m *Manager) deregister(listener *Listener) {
	delete(m.uids, listener.Uid)
	m.bcChannel.Unregister(listener.Channel)
	close(listener.Channel)
}

func (m *Manager) deleteBroadcast() {
	m.bcChannel.Close()
}

func (m *Manager) OpenListener(uid string) chan interface{} {
	listener := make(chan interface{})
	m.open <- &Listener{
		Uid:     uid,
		Channel: listener,
	}
	return listener
}

func (m *Manager) CloseListener(uid string, channel chan interface{}) {
	m.close <- &Listener{
		Uid:     uid,
		Channel: channel,
	}
}

func (m *Manager) DeleteBroadcast() {
	m.delete <- struct{}{}
}

func (m *Manager) SubmitUserInfo(text string) {
	msg := &Message{
		Flag: Msg_Flag_User_info,
		Text: text,
	}
	m.messages <- msg
}

func (m *Manager) TimeProvider() {
	for x := range time.Tick(3 * time.Second) {
		m.SubmitUserInfo(x.String())
	}
}
