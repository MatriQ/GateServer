package Handlers

var EnterThreadHandler struct {
	message *Gate.Messager
}

func (hander *UserLoginHander) GetMessage() *Gate.Messager {
	return hander.message
}

func (handler *EnterThreadHandler) Action() {
	var msg = hander.GetMessage()
}
