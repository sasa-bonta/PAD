package main

import (
	"PAD1/common"
	"PAD1/subscriber/cmd"
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

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Topic: ")
	topic, _ := reader.ReadString('\n')
	topic = strings.ToLower(strings.TrimSpace(topic))

	messageObj := &common.Message{Action: common.SUBSCRIBE, Topic: topic}
	messageJson, _ := json.Marshal(messageObj)
	messageToSend := string(messageJson) + "\n"

	cmd.WriteToConnection(conn, []byte(messageToSend))
	message, _ := bufio.NewReader(conn).ReadString('\n')
	log.Print("Server relay: " + message)

	handleSub(conn)
}

func handleSub(conn net.Conn) {
	messageJson, _ := bufio.NewReader(conn).ReadString('\n')
	var message common.Message
	cmd.UnmarshalJsonToMessage(messageJson, &message)
	log.Print("New message: " + message.Text)
	handleSub(conn)
}
