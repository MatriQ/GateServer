package Gate

type MessageManager struct {
	Handers  [int]interface{}
	Messages [int]interface{}
}

func NewMessagePool() {
	var ret = &MessageManager{}
	return ret
}

func (*MessageManager) Register(id int, handler interface{}, message interface{}) {
}

func (*MessageManager) CreateHandler(id int) interface{} {
}

func (*MessageManager) CreateMessage(id int) interface{} {
}
