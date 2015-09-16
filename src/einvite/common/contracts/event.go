package contracts

import ()

type Event struct {
	Id   string
	Name string

	MainActivity int
	Activities   []*Activity

	Creator      *User
	Participants []*Participant

	Version int
}
