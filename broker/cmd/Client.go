package cmd

import "net"

type Client struct {
	Id         string
	Connection net.Conn
}
