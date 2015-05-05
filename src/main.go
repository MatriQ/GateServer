package main

import (
	"GateServer"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	_server := GateServer.GetServer()
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
		case "":
		default:
			fmt.Printf("Unknow command:%s\n command help to get help info\n", command)
		}
	}
}
