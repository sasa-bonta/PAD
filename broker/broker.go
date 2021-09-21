package main

import (
	"PAD1/broker/cmd"
	"PAD1/common"
	"bufio"
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

		// Adding new client to the list of clients
		ps.AddClient(&client)

		fmt.Println("Client connected.")

		// Print client connection address.
		fmt.Println("Client " + conn.RemoteAddr().String() + " connected.")

		// Handle connections concurrently in a new goroutine.
		go handleConnection(conn)
	}
}

//func getMessage(conn, net.Conn)  {
//
//}

// handleConnection handles logic for a single connection request.
func handleConnection(conn net.Conn) {
	// Buffer client input until a newline.
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	// Close left clients.
	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	// Print response message, stripping newline character.
	log.Println("Client message:", string(buffer[:len(buffer)-1]))

	// Send response message to the client.
	conn.Write(buffer)

	// Restart the process.
	handleConnection(conn)
}

