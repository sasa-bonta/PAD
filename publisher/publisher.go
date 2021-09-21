package main

import (
	"PAD1/common"
	"PAD1/publisher/cmd"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting " + common.ConnectionType + " server on " + common.HostPort)
	conn, err := net.Dial(common.ConnectionType, common.HostPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	handlePub(conn)
}

func handlePub(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Topic: ")
	topic, _ := reader.ReadString('\n')
	topic = strings.ToLower(strings.TrimSpace(topic))
	fmt.Print("Message: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	messageObj := &common.Message{Action: common.PUBLISH, Topic: topic, Text: text}
	messageJson, _ := json.Marshal(messageObj)
	messageToSend := string(messageJson) + "\n"

	cmd.WriteToConnection(conn, []byte(messageToSend))
	message, _ := bufio.NewReader(conn).ReadString('\n')
	log.Print("Server relay: " + message)

	handlePub(conn)
}
