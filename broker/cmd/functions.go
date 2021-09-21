package cmd

import (
	"PAD1/common"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

func GetBuffer(conn net.Conn) []byte {
	// Buffer client input until a newline.
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	// Close left clients.
	if err != nil {
		fmt.Println("Client left.")
		CloseConnection(conn)
		return nil
	}
	return buffer
}

func DecodeBuffer(buffer []byte) string {
	if len(buffer) < 2 {
		fmt.Println("Cannot decode empty buffer")
	}
	return string(buffer[:len(buffer)-1])
}

func UnmarshalBufferToMessage(buffer []byte, message *common.Message) string {
	messageJson := DecodeBuffer(buffer)
	err := json.Unmarshal([]byte(messageJson), message)
	if err != nil {
		fmt.Println("Cannot unmarshal json")
	}
	return messageJson
}

func CloseConnection(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		fmt.Println("Cannot close connection")
	}
}

func CloseListener(conn net.Listener) {
	err := conn.Close()
	if err != nil {
		fmt.Println("Cannot close listener")
	}
}

func WriteToConnection(conn net.Conn, buffer []byte) {
	_, err := conn.Write(buffer)
	if err != nil {
		fmt.Println("Cannot write to connection")
	}
}
