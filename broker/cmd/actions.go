package cmd

import (
	"fmt"
	"net"
)

func (ps *ClientsList) AddClient(client *Client, conn net.Conn) *ClientsList {
	ps.Clients = append(ps.Clients, *client)
	fmt.Println("New client address:", conn.RemoteAddr().String(), " id:", client.Id, " total:", len(ps.Clients))
	return ps
}

func (ps *ClientsList) Subscribe(client *Client, topic string) *ClientsList {
	clientSubs := ps.GetSubscriptions(topic, client)

	if len(clientSubs) > 0 {
		return ps
	}

	newSubscription := Subscription{
		Topic:  topic,
		Client: client,
	}

	ps.Subscriptions = append(ps.Subscriptions, newSubscription)
	return ps
}

func (ps *ClientsList) Publish(topic string, message string) {
	subscriptions := ps.GetSubscriptions(topic, nil)

	for _, sub := range subscriptions {
		fmt.Printf("Sending to client id %s message is %s \n", sub.Client.Id, message)
		message = message + "\n"
		sub.Client.Connection.Write([]byte(message))
	}
}

func (ps *ClientsList) GetSubscriptions(topic string, client *Client) []Subscription {

	var subscriptionList []Subscription

	for _, subscription := range ps.Subscriptions {

		if client != nil {

			if subscription.Client.Id == client.Id && subscription.Topic == topic {
				subscriptionList = append(subscriptionList, subscription)

			}
		} else {

			if subscription.Topic == topic {
				subscriptionList = append(subscriptionList, subscription)
			}
		}
	}

	return subscriptionList
}
