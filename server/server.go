package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/vSterlin/tcp/util"
)

type Server struct {
	Connections *[]net.Conn
	IP          string
	Port        int
}

func (s *Server) Listen() {

	addr := util.BuildAddress(s.IP, s.Port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error: Was not able to listen for connection.")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error: Wasn't able to accept connection.")
			continue
		}
		
		(*s.Connections) = append((*s.Connections), conn)
		go s.handleClient(conn)
	}
}

func (s *Server) handleClient(conn net.Conn) {
	// defer conn.Close()
	clientReader := bufio.NewReader(conn)
	for {

		clientRequestText, err := clientReader.ReadString('\n')
		if err != nil {
			fmt.Println("An error has occurred between client and server.")
			break
		}
		// remove \n
		clientRequestText = util.TrimString(clientRequestText)

		addr := conn.RemoteAddr().String()
		splitAddr := strings.Split(addr, ":")
		ip, port := splitAddr[0], splitAddr[1]

		message := "\nMessage received from %s\nSender’s Port: %s\nMessage: \"%s\"\n"
		fmt.Printf(message, ip, port, clientRequestText)
		fmt.Print(">>> ")
		// Responding to the client request
		// conn.Write([]byte("Got it lol\n"))

	}
}
