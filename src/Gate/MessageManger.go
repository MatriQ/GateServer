package Gate

import (
	"Gate/NetSession"
	"reflect"
)

type Messager interface {
	GetSession() *NetSession.NetSession
}

type Hander interface {
	Action()
	GetMessage() *Messager
}

type MessageManager struct {
	Handers  map[int]reflect.Type
	Messages map[int]reflect.Type
}

func NewMessagePool() {
	var ret = &MessageManager{
		Handers:   make(map[int]reflect.Type),
		Messagers: make(map[int]reflect.Type),
	}
	return ret
}

func (manager *MessageManager) register(id int, handlerType reflect.Type, messageType reflect.Type) {
	manager.Handers[id] = handlerType
	manager.Messages[id] = messageType
}

func (manager *MessageManager) CreateHandler(id int) *Hander {
	var t = manager.Handers[id]
	if t == nil {
		//error
	}

	var ret = &reflect.New(t)
	return ret
}

func (manager *MessageManager) CreateMessage(id int) *Messager {
	var t = manager.Handers[id]
	if t == nil {
		//error
	}
	return &reflect.New(t)
}
