package main

import (
	"PAD1/common"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Start the client and connect to the server.
	fmt.Println("Starting " + common.ConnectionType + " server on " + common.HostPort)
	conn, err := net.Dial(common.ConnectionType, common.HostPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	handleSub(conn)
}

func handleSub(conn net.Conn) {
	// Create new reader from Stdin.
	reader := bufio.NewReader(os.Stdin)
	// Prompting message.
	fmt.Print("Topic: ")
	// Read in input until newline, Enter key.
	topic, _ := reader.ReadString('\n') // Prompting message.
	topic = strings.ToLower(strings.TrimSpace(topic))
	fmt.Print("Message: ")
	// Read in input until newline, Enter key.
	text, _ := reader.ReadString('\n')

	messageTopic := &common.Message{Action: common.PUBLISH, Topic: topic, Text: strings.TrimSpace(text)}
	messageToSend, _ := json.Marshal(messageTopic)
	messageJson := string(messageToSend) + "\n"

	// Send to socket connection.
	conn.Write([]byte(messageJson))
	// Listen for relay.
	message, _ := bufio.NewReader(conn).ReadString('\n')
	// Print server relay.
	log.Print("Server relay: " + message)

	handleSub(conn)
}
