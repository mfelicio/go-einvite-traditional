package contracts

import ()

type Activity struct {
	Id int

	Type ActivityType

	InterestedParticipants []int

	Where int
	What  int
	When  int

	Wheres []*Choosable
	Whats  []*Choosable
	Whens  []*Choosable
}
