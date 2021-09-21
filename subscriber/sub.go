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

	// Create new reader from Stdin.
	reader := bufio.NewReader(os.Stdin)
	// Prompting message.
	fmt.Print("Topic: ")
	// Read in input until newline, Enter key.
	topic, _ := reader.ReadString('\n') // Prompting message.
	topic = strings.ToLower(strings.TrimSpace(topic))

	messageTopic := &common.Message{Action: common.SUBSCRIBE, Topic: topic}
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

// run loop forever, until exit.
func handleSub(conn net.Conn) {
	// Listen for relay.
	message, _ := bufio.NewReader(conn).ReadString('\n')

	// Print server relay.
	log.Print("New message: " + message)
	handleSub(conn)
}
