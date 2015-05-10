package Gate

import (
	"Gate/NetSession"
)

type Messager interface {
	GetSession() *NetSession.NetSession
}

type Hander interface {
	Action()
	GetMessage() *Messager
}

type MessageManager struct {
	Handers  [int]interface{}
	Messages [int]interface{}
}

func NewMessagePool() {
	var ret = &MessageManager{}
	return ret
}

func (manager *MessageManager) register(id int, handler *Hander, message *Messager) {
	manager.Handers[id] = handler
	manager.Messages[id] = message
}

func (manager *MessageManager) CreateHandler(id int) *Hander {
	var ret = &Hander{}
	return ret
}

func (manager *MessageManager) CreateMessage(id int) *Messager {
	var ret = &Messager{}
	return ret
}
