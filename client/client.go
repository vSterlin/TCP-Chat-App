package client

import (
	"fmt"
	"net"

	"github.com/vSterlin/tcp/utils"
)

type Client struct {
	Connections *[]net.Conn
}

func (c *Client) Connect(ip string, port int) {

	addr := utils.BuildAddress(ip, port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	(*c.Connections) = append((*c.Connections), conn)

}

func (c *Client) SendMessage(id int, input string) {
	// clientReader := bufio.NewReader(os.Stdin)
	conn := (*c.Connections)[id-1]
	// serverReader := bufio.NewReader(conn)

	if input[len(input)-1] != '\n' {
		input += "\n"
	}

	conn.Write([]byte(input))
	// serverRes, _ := serverReader.ReadString('\n')
	// fmt.Println(serverRes)

}

func (c *Client) CloseConnections() {
	for _, conn := range *c.Connections {
		conn.Close()
	}
}
