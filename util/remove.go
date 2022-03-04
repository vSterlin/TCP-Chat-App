package util

import "net"

func RemoveIndex(slice *[]net.Conn, index int) {
	newSlice := append((*slice)[:index], (*slice)[index+1:]...)
	*slice = newSlice
}
