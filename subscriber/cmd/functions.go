package cmd

import (
	"PAD1/common"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func UnmarshalJsonToMessage(messageJson string, message *common.Message) string {
	err := json.Unmarshal([]byte(messageJson), message)
	if err != nil {
		fmt.Println("Cannot unmarshal json")
		os.Exit(1)
	}
	return messageJson
}

func WriteToConnection(conn net.Conn, buffer []byte) {
	_, err := conn.Write(buffer)
	if err != nil {
		fmt.Println("Cannot write to connection")
		os.Exit(1)
	}
}

func CloseConnection(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		fmt.Println("Cannot close connection")
		os.Exit(1)
	}
}
