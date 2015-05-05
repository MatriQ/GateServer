package user

import (
	"NetSession"
	"fmt"
)

type UserManager struct {
}

func NewUserManager() *UserManager {
	return &UserManager{}
}
func RegisterUser(session *NetSession, string userName, string userPass) {

}

func UserLogin(session *NetSession, string userName, string userPass) {

}

func AgetnLogin(session *NetSession, mst interface{}) {

}
