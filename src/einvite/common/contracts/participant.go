package contracts

import ()

type Participant struct {
	Id     int
	Email  string
	User   *User
	Role   ParticipantRole
	Status ParticipantStatus
}
