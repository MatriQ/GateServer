package Handlers

import (
	"Gate"
	"Gate/Manager"
)

type UserLoginHander struct {
	message *Gate.Messager
}

func (hander *UserLoginHander) GetMessage() *Gate.Messager {
	return hander.message
}

func (hander *UserLoginHander) Action() {
	var msg = hander.GetMessage()
	Manager.GetInstance().UserManager.UserLogin(msg.GetSession(), msg.GetUserName(), msg.GetUserPass())
}
