package Handlers

import (
	"Gate/Manager"
)

type UserLoginHander struct {
}

func (hander *UserLoginHander) GetMessage() {
	return nil
}

func (hander *UserLoginHander) Action() {
	var msg = hander.GetMessage()
	Manager.GetInstance().UserManager.UserLogin(msg.GetSession(), msg.GetUserName(), msg.GetUserPass())
}
