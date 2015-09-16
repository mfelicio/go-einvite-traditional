package services

import (
	"einvite/common/contracts"
)

type UserService interface {
	List() ([]*contracts.User, error)
	Create(userDto *contracts.User) (*contracts.User, error)
	SaveWithCredentials(userDto *contracts.User, credentialsDto *contracts.UserAuthCredentials) (*contracts.User, error)
	Get(id string) (*contracts.User, error)
}

type EventService interface {
	CreateEvent(event *contracts.Event) (*contracts.Event, error)
}
