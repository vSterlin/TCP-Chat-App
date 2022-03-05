package app

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/vSterlin/tcp/client"
	"github.com/vSterlin/tcp/server"
	"github.com/vSterlin/tcp/util"
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
		fmt.Println("Options to choose:")
		fmt.Println("help")
		fmt.Println("myip")
		fmt.Println("myport")
		fmt.Println("connect <dest> <port #>")
		fmt.Println("list")
		fmt.Println("terminate <connection ID>")
		fmt.Println("send <connection ID> <message>")
		fmt.Println("exit")
		fmt.Println("\nPlease type one of the command (case sensitive):")
		
		
		input := a.ReadUserInput()
		fmt.Println("")

		splitInput := strings.Split(input, " ")
		firstKeyword := splitInput[0]
		switch firstKeyword {
			case "help":
				fmt.Println("\nMyIP:")
				fmt.Println("Display the IP address of this process.")
				fmt.Println("\nMyPort:")
				fmt.Println("Display the port on which this process is")
				fmt.Println("listening for incoming connections.")
				fmt.Println("\nConnect <dest> <port #>:")
				fmt.Println("Establish a TCP connection to <dest> at <port #>.")
				fmt.Println("\nList:")
				fmt.Println("Display the list of connections this process is part of.")
				fmt.Println("\nTerminate <ID> : ")
				fmt.Println("Terminated connection listed under ID from List.")
				fmt.Println("\nSend <connection ID> <message>:")
				fmt.Println("Send the message to the IP via connection ID.")
				fmt.Println("\nExit:")
				fmt.Println("Close all connections and terminate the process.")
				fmt.Println("")
			
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
