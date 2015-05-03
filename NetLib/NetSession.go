package NetSession

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
	"time"
)

//最大连接数
const SESSION_MAX = 20000

var locker sync.Mutex

type NetSession struct {
	ID          int
	Conn        net.Conn
	StartTime   time.Time
	LastMsgTime time.Time
	buff        *bytes.Buffer
	readBuff    []byte
	buffOffset  int
	manager     *SessionManager
	breakChan   chan bool
}

func (session *NetSession) String() string {
	return "[Session:" + strconv.Itoa(session.ID) + ",StartTime:" +
		session.StartTime.Format("2006-01-02 15:04:05") +
		",LastMsgTime:" + session.LastMsgTime.Format("2006-01-02 15:04:05")
}

func (session *NetSession) connectionHandle() {
	conn := session.Conn
	for {
		bufflen, err := conn.Read(session.readBuff)
		if err != nil {
			if err == io.EOF {
				close(session.breakChan)
				session.RemveFromManager()
				fmt.Println("session:", session.ID, "read error：", err)
			}
			//return
		}
		select {
		case <-session.breakChan:
			conn.Close()
			return
		default:
		}
		if err == nil {
			session.LastMsgTime = time.Now()
			fmt.Printf("session:%d[%s] recived msg:%s\n", session.ID, session.Conn.RemoteAddr(), string(session.readBuff[:bufflen]))
		}
	}
}
func (session *NetSession) RemveFromManager() {
	session.manager.RemoveSession(session.ID)
}
func (session *NetSession) SendMsg(msg string) {
	msgBuff := []byte(msg)
	//msglen := len(msgBuff)
	/*if session.buffOffset+msglen < 1024 {
		session.buff.Write(msgBuff)
	}*/
	_, err := session.Conn.Write(msgBuff[0:])
	if err != nil {
		fmt.Println("NetSession.SendMsg error:", err)
	}
}

//****************SessionManager*******************//
type SessionManager struct {
	freeIDs  []int
	sessions map[int]*NetSession
}

func NewSessionManager() *SessionManager {
	manager := &SessionManager{make([]int, SESSION_MAX), make(map[int]*NetSession)}
	for i := 0; i < SESSION_MAX; i++ {
		manager.freeIDs[i] = 100000 + i
	}
	return manager
}

func (sessionManager *SessionManager) AddSession(conn *net.Conn) (err error) {
	locker.Lock()
	defer locker.Unlock()
	if len(sessionManager.freeIDs) == 0 {
		err = errors.New("SessionManager is full")
		return
	}
	sessionID := sessionManager.freeIDs[0]
	sessionManager.freeIDs = sessionManager.freeIDs[1:]
	session := &NetSession{
		ID:          sessionID,
		Conn:        *conn,
		LastMsgTime: time.Now(),
		StartTime:   time.Now(),
		buff:        nil,
		readBuff:    make([]byte, 1024),
		buffOffset:  0,
		manager:     sessionManager,
		breakChan:   make(chan bool),
	}
	go session.connectionHandle()
	session.buff = bytes.NewBuffer(make([]byte, 0, 1024))
	sessionManager.sessions[sessionID] = session
	return
}
func (manager *SessionManager) RemoveSession(id int) {
	locker.Lock()
	defer locker.Unlock()
	manager.freeIDs = append(manager.freeIDs, id)
	delete(manager.sessions, id)
}
func (manager *SessionManager) Close() {
	for _, session := range manager.sessions {
		session.Conn.Close()
	}
	manager.sessions = make(map[int]*NetSession)
}
func (manager *SessionManager) GetSessionCount() int {
	return len(manager.sessions)
}
func (manager *SessionManager) GetSession(id int) *NetSession {
	if session, ok := manager.sessions[id]; ok {
		return session
	}
	return nil
}
