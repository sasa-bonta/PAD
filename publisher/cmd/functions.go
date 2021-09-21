package cmd

import (
	"fmt"
	"net"
)

func WriteToConnection(conn net.Conn, buffer []byte) {
	_, err := conn.Write(buffer)
	if err != nil {
		fmt.Println("Cannot write to connection")
	}
}
