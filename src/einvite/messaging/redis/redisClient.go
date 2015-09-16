package redis

import (
	"einvite/messaging"
	//"strings"
)

//Implements Messaging.Message
type redisClient struct {
}

func (this *redisClient) Open(handler messaging.MessageHandler) {

}

func (this *redisClient) Close() {

}

func (this *redisClient) Broadcast(message messaging.Message) {

}

func NewMessagingClient() messaging.MessagingClient {

	return &redisClient{}
}
