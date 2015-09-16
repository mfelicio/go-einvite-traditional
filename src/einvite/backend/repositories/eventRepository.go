package repositories

import (
	"einvite/common/contracts"
)

type EventRepository interface {
	Create(eventDto *contracts.Event) (*contracts.Event, error)
}
