package cmd

import (
	"PAD1/common"
	"encoding/json"
	"fmt"
	"net"
)

func UnmarshalJsonToMessage(messageJson string, message *common.Message) string {
	err := json.Unmarshal([]byte(messageJson), message)
	if err != nil {
		fmt.Println("Cannot unmarshal json")
	}
	return messageJson
}

func WriteToConnection(conn net.Conn, buffer []byte) {
	_, err := conn.Write(buffer)
	if err != nil {
		fmt.Println("Cannot write to connection")
	}
}
