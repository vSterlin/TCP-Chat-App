package app

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/vSterlin/chat/client"
	"github.com/vSterlin/chat/server"
	"github.com/vSterlin/chat/util"
)

type App struct {
	Client      *client.Client
	Server      *server.Server
	Connections *[]net.Conn
}

func NewApp() *App {

	// gets port number from cli argument
	port, _ := strconv.Atoi(os.Args[1])

	// array to store all connections
	netConns := &[]net.Conn{}
	s := &server.Server{Connections: netConns, IP: util.GetIp(), Port: port}
	c := &client.Client{Connections: netConns}

	return &App{Client: c, Server: s, Connections: netConns}
}

func (a *App) Run() {
	go a.Server.Listen()

	listUsed := false
	for {
		input := a.ReadUserInput()
		fmt.Println(input)

		splitInput := strings.Split(input, " ")
		firstKeyword := splitInput[0]
		switch firstKeyword {
		case "connect":
			ip := splitInput[1]
			port, _ := strconv.Atoi(splitInput[2])
			a.Client.Connect(ip, port)
		case "send":
			id, _ := strconv.Atoi(splitInput[1])

			// adds 3rd and next split words to rebuild message
			message := strings.Join(splitInput[2:], " ")
			a.Client.SendMessage(id, message)
		case "myip":
			fmt.Println(a.Server.IP)
		case "myport":
			fmt.Println(a.Server.Port)
		case "list":
			a.ListConnections()
			listUsed = true
		case "terminate":
			if listUsed == true {
				indexToTerminate, _ := strconv.Atoi(splitInput[1])
				a.TerminateConnection(indexToTerminate)
				listUsed = false
			} else {
				fmt.Println("Error: You did not use List first. How would you know which IP to terminate?")
			}

		case "exit":
			os.Exit(0)

		}

	}

}

func (a *App) ListConnections() {

	for i, conn := range *a.Connections {
		// using index as id
		fmt.Printf("%d: %s\n", i+1, conn.RemoteAddr().String())
	}
}

func (a *App) ReadUserInput() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(">>> ")
	str, _ := reader.ReadString('\n')

	fmt.Println([]byte(str))

	str = util.TrimString(str)
	return str

}

func (a *App) CloseConnections() {
	for _, conn := range *a.Connections {
		conn.Close()
	}
}

func (a *App) TerminateConnection(i int) {
	conn := (*a.Connections)[i-1]
	conn.Close()
	util.RemoveIndex(a.Connections, i-1)

}
