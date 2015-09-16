package messaging

import ()

type MessagingClient interface {
	Open(handler MessageHandler)
	Close()

	Broadcast(message Message)
}

type MessageHandler func(Message)
