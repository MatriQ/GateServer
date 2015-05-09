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
func RegisterUser(session *NetSession, userName string, userPass string) {

}

func UserLogin(session *NetSession, userName string, userPass string) {

}

func AgetnLogin(session *NetSession, mst interface{}) {

}
