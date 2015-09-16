package services

import (
	"einvite/backend/repositories"
	"einvite/common/contracts"
	"einvite/common/services"
)

type eventService struct {
	eventRepository repositories.EventRepository
}

func (this eventService) CreateEvent(event *contracts.Event) (*contracts.Event, error) {

	//creates new event on the database

	newEvent, err := this.eventRepository.Create(event)

	//sends invites to every participant
	//TODO: send app messages to participants who are already users

	return newEvent, err
}

func NewEventService(eventRepository repositories.EventRepository) services.EventService {

	return &eventService{eventRepository}
}
