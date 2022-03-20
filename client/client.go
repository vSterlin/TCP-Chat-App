package client

import (
	"fmt"
	"net"

	"github.com/vSterlin/chat/util"
)

type Client struct {
	Connections *[]net.Conn
}

func (c *Client) Connect(ip string, port int) {

	addr := util.BuildAddress(ip, port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Error: Invalid connection, try again.")
		return
	}
	fmt.Println("Connection has been established!")
	fmt.Println("")
	(*c.Connections) = append((*c.Connections), conn)

}

func (c *Client) SendMessage(id int, input string) {
	// clientReader := bufio.NewReader(os.Stdin)
	conn := (*c.Connections)[id-1]
	// serverReader := bufio.NewReader(conn)

	if input[len(input)-1] != '\n' {
		input += "\n"
	}

	_, err := conn.Write([]byte(input))
	if err != nil {
		fmt.Println("Error: Was not able to send message.")
	}
	// serverRes, _ := serverReader.ReadString('\n')
	// fmt.Println(serverRes)

}

func (c *Client) CloseConnections() {
	for _, conn := range *c.Connections {
		conn.Close()
	}
}
