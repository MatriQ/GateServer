package Manager

import (
	"User"
	"fmt"
	"sync"
)

type ManagerPool struct {
	ServerManager interface{}
	UserManager   *user.UserManager
}

var instance *ManagerPool
var locker sync.Locker

func GetInstance() *ManagerPool {
	if instance == nil {
		locker.Lock()
		if instance == nil {
			instance = &ManagerPool{
				ServerManager: nil,
				UserManager:   User.NewUserNamager(),
			}
		}
		locker.Unlock()
	}
}
