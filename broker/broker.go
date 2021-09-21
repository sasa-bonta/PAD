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
	fmt.Println("Starting " + common.ConnectionType + " server on " + common.HostPort)
	l, err := net.Listen(common.ConnectionType, common.HostPort)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer func(l net.Listener) {
		cmd.CloseListener(l)
	}(l)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}

		client := cmd.Client{
			Id:         autoId(),
			Connection: conn,
		}

		fmt.Println("Client " + conn.RemoteAddr().String() + " connected.")
		ps.AddClient(&client, conn)

		buffer := cmd.GetBuffer(conn)
		var message common.Message
		messageJson := cmd.UnmarshalBufferToMessage(buffer, &message)

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

		log.Println("Client message:", messageJson)
		cmd.WriteToConnection(conn, buffer)
	}
}

func handlePublisher(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		fmt.Println("Client left.")
		cmd.CloseConnection(conn)
	}

	var message common.Message
	cmd.UnmarshalBufferToMessage(buffer, &message)

	fmt.Println("This is publish new message")
	ps.Publish(message.Topic, message.Text)
	cmd.WriteToConnection(conn, buffer)
	handlePublisher(conn)
}
