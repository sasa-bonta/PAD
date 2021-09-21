package main

import (
	"PAD1/broker/cmd"
	"PAD1/common"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"log"
	"net"
	"os"
)

func autoId() string {
	return uuid.Must(uuid.NewV4(), nil).String()
}

var ps = &cmd.ClientsList{}

func main() {
	// Start the server and listen for incoming connections.
	fmt.Println("Starting " + common.ConnectionType + " server on " + common.HostPort)
	l, err := net.Listen(common.ConnectionType, common.HostPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()

	// run loop forever, until exit.
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}

		client := cmd.Client{
			Id:         autoId(),
			Connection: conn,
		}

		fmt.Println("---------------------------------------")
		// Print client connection address.
		fmt.Println("Client " + conn.RemoteAddr().String() + " connected.")
		// Adding new client to the list of clients
		ps.AddClient(&client, conn)

		// Buffer client input until a newline.
		buffer := getBuffer(conn)
		messageJson := decodeBuffer(buffer)
		var message common.Message
		json.Unmarshal([]byte(messageJson), &message)

		// Verify type of client
		switch message.Action {

		case common.SUBSCRIBE:
			ps.Subscribe(&client, message.Topic)
			fmt.Println("new subscriber to topic", message.Topic, len(ps.Subscriptions), client.Id)
			break

		case common.PUBLISH:
			fmt.Println("This is publish new message")
			ps.Publish(message.Topic, message.Text)
			go handlePublisher(conn)
			break

		default:
			break
		}

		// Print response message, stripping newline character.
		log.Println("Client message:", messageJson)
		// Send response message to the client.
		conn.Write(buffer)
	}
}

func handlePublisher(conn net.Conn) {
	// Buffer client input until a newline.
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	// Close left clients.
	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
	}

	messageJson := decodeBuffer(buffer)
	var message common.Message
	json.Unmarshal([]byte(messageJson), &message)

	fmt.Println("This is publish new message")
	ps.Publish(message.Topic, message.Text)
	conn.Write(buffer)
	handlePublisher(conn)
}

func getBuffer(conn net.Conn) []byte {
	// Buffer client input until a newline.
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	// Close left clients.
	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return nil
	}
	return buffer
}

func decodeBuffer(buffer []byte) string {
	return string(buffer[:len(buffer)-1])
}
