package util

import (
	"net"
	"log"
	"strconv"
	"strings"
)

func GetIp() string {
	conn, issue := net.Dial("udp", "8.8.8.8:80")
	if issue != nil {
		log.Fatal("Failed to get IP: ", err)
	}
	defer conn.Close()

	// get ip address without port number
	addr := strings.Split(conn.LocalAddr().String(), ":")[0]
	return addr
}

func BuildAddress(ip string, port int) string {
	addr := ip + ":" + strconv.Itoa(port)
	return addr
}
