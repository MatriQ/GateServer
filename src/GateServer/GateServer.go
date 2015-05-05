package GateServer

import (
	"NetSession"
	"fmt"
	"net"
	"os"
	//"reflect"
	"regexp"
	"strconv"
	"sync"
	//"time"
	"bufio"
)

const (
	Port = 8092
)

/*
Server status
*/
const (
	Closed = iota
	Starting
	Running
	Closing
)

/*
Gate server define
*/
type Server struct {
	Version        string
	SessionManager *NetSession.SessionManager
	State          int
	Listener       net.Listener
	closeChan      chan bool
}

/*
start server
*/
func (server *Server) Start() {
	if server.State != Closed {
		fmt.Printf("Close server failed,the state of server is %d", server.State)
		return
	}
	once.Do(func() {
		server.State = Starting
		fmt.Println("Start Listen:", Port)
		listener, err := net.Listen("tcp", ":"+strconv.Itoa(Port))
		server.Listener = listener
		if err != nil {
			fmt.Println("server start failed,", err)
			return
		}
		fmt.Println("Ready accept")
		server.State = Starting
		for {
			select {
			case <-server.closeChan:
				server.State = Closing
				return
			default:
			}
			fmt.Println("Ready accept in loop ", listener)
			if conn, err1 := listener.Accept(); err1 != nil {
				fmt.Println("Accept Listener:", err1)
			} else {
				fmt.Println("Accept Listener:", conn)
				server.SessionManager.AddSession(&conn)
			}
		}
		server.State = Closed
	})
}

/*
stop server
*/
func (server *Server) Stop() {
	if server.State != Running {
		fmt.Printf("Stop server failed,the state of server is %d", server.State)
	} else {
		close(server.closeChan)
		server.SessionManager.Close()
		server.Listener.Close()
		server.Listener = nil
		fmt.Println("Server Stoped")
	}
	server.State = Closed
}

/*
get singleinstance server
*/
func GetServer() *Server {
	if _instance == nil {
		_instance = &Server{
			Version:        "0.0.1",
			SessionManager: NetSession.NewSessionManager(),
			State:          Closed,
			Listener:       nil,
			closeChan:      make(chan bool)}
	}
	return _instance
}
func (server *Server) GetAllServerSession() []*NetSession.NetSession {
	return server.SessionManager.GetAllSessions()
}
func (server *Server) String() {
	fmt.Println("Server:", server.Version, server.SessionManager)
}

var once sync.Once
var _instance *Server

func main() {
	_server := GetServer()
	running := true
	for running {
		fmt.Println("Please input command:")
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		args := string(data)
		reg := regexp.MustCompile("\\S+")
		command := reg.FindAllString(args, -1)
		fmt.Println(command)
		//fmt.Println("Got command:", command)
		if len(command) == 0 {
			continue
		}
		switch command[0] {
		case "quit":
			fmt.Println("server quit")
			running = false
		case "start":
			go _server.Start()
		case "stop":
			go _server.Stop()
		case "state":
			fmt.Println("Server State:", _server)
		case "getsessioncount":
			fmt.Println("Session count:", _server.SessionManager.GetSessionCount())
		case "getsession":
			if len(command) != 2 {
				fmt.Println("Command getsession less args\n\tget session <id>")
				continue
			}
			id, err := strconv.Atoi(command[1])
			if err != nil {
				fmt.Println("Command getsession args 1 is not a integer")
				continue
			}
			session := _server.SessionManager.GetSession(id)
			fmt.Println("get session info:", session)
		case "sendmsg":
			if len(command) != 3 {
				fmt.Println("Command sendmsg less args\n\tget session <id> <message>")
				continue
			}
			id, err := strconv.Atoi(command[1])
			if err != nil {
				fmt.Println("Command sendmsg args 1 is not a integer")
				continue
			}
			session := _server.SessionManager.GetSession(id)
			if session != nil {
				session.SendMsg(command[2])
			}
		case "help":
			fmt.Println("GateServer Command:\n\tquit:\tquit server\n\tstart:\tstart server\n\tstop:\tstop server\n\tstate:\tshow server infos\n\thelp:\tto get help")
		default:
			fmt.Printf("Unknow command:%s\n command help to get help info\n", command)
		}
	}
}
