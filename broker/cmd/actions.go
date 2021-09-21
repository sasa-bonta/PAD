package cmd

import (
	"fmt"
)

func (ps *ClientsList) AddClient(client *Client) *ClientsList {
	ps.Clients = append(ps.Clients, *client)
	fmt.Println("adding new client to the list", client.Id, len(ps.Clients))
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
		fmt.Println(sub.Client.Connection)
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

//func (ps *ClientsList) RemoveClient(client Client) *ClientsList {
//
//	// first remove all subscriptions by this client
//
//	for index, sub := range ps.Subscriptions {
//
//		if client.Id == sub.Client.Id {
//			ps.Subscriptions = append(ps.Subscriptions[:index], ps.Subscriptions[index+1:]...)
//		}
//	}
//
//	// remove client from the list
//
//	for index, c := range ps.Clients {
//
//		if c.Id == client.Id {
//			ps.Clients = append(ps.Clients[:index], ps.Clients[index+1:]...)
//		}
//
//	}
//
//	return ps
//}
//
//
//
//func (client *Client) Send(message []byte) error {
//
//	return client.Connection.WriteMessage(1, message)
//
//}
//
//func (ps *ClientsList) Unsubscribe(client *Client, topic string) *ClientsList {
//
//	//clientSubscriptions := ps.GetSubscriptions(topic, client)
//	for index, sub := range ps.Subscriptions {
//
//		if sub.Client.Id == client.Id && sub.Topic == topic {
//			// found this subscription from client and we do need remove it
//			ps.Subscriptions = append(ps.Subscriptions[:index], ps.Subscriptions[index+1:]...)
//		}
//	}
//
//	return ps
//
//}
//
//func (ps *ClientsList) HandleReceiveMessage(client Client, messageType int, payload []byte) *ClientsList {
//
//	m := common.Message{}
//
//	err := json.Unmarshal(payload, &m)
//	if err != nil {
//		fmt.Println("This is not correct message payload")
//		return ps
//	}
//
//	switch m.Action {
//
//	case common.PUBLISH:
//
//		fmt.Println("This is publish new message")
//
//		ps.Publish(m.Topic, m.Text, nil)
//
//		break
//
//	case common.SUBSCRIBE:
//
//		ps.Subscribe(&client, m.Topic)
//
//		fmt.Println("new subscriber to topic", m.Topic, len(ps.Subscriptions), client.Id)
//
//		break
//
//	default:
//		break
//	}
//
//	return ps
//}
