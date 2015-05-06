package GameServer

import (
	"NetSession"
	//"fmt"
)

type ServerManager struct {
}

/*
register gameserver
*/
func (manager *ServerManager) RegisterServer(session *NetSession.NetSession, server interface{}) {

}

/*
send all server to client
*/
func (manager *ServerManager) SendServerList(session *NetSession.NetSession) {

}
