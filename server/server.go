package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/vSterlin/chat/util"
)

type Server struct {
	Connections *[]net.Conn
	IP          string
	Port        int
}

func (s *Server) Listen() {

	addr := util.BuildAddress(s.IP, s.Port)
	ln, _ := net.Listen("tcp", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
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
