package cmd

import (
	"fmt"
	"net"
	"os"
)

func WriteToConnection(conn net.Conn, buffer []byte) {
	_, err := conn.Write(buffer)
	if err != nil {
		fmt.Println("Cannot write to connection")
		os.Exit(1)
	}
}
