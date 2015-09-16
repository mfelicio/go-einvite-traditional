package contracts

type ParticipantRole int

const (
	ParticipantRole_Owner       ParticipantRole = 1
	ParticipantRole_Contributor ParticipantRole = 2
	ParticipantRole_Guest       ParticipantRole = 4
)

type ParticipantStatus int

const (
	ParticipantStatus_PendingAccept ParticipantStatus = 1
	ParticipantStatus_Accepted      ParticipantStatus = 2
	ParticipantStatus_Rejected      ParticipantStatus = 4
)

//Example: Watch a movie, Have dinner, Go drinking, Go walk, Trip
type ActivityType int

const (
	ActivityType_Movie  ActivityType = 1
	ActivityType_Food   ActivityType = 2
	ActivityType_Travel ActivityType = 3
)

type ChoosableType int

const (
	ChoosableType_Where ChoosableType = 1
	ChoosableType_What  ChoosableType = 2
	ChoosableType_When  ChoosableType = 3
)

type ChoosableDataType int

const (
	ChoosableDataType_FreeText  ChoosableDataType = 1
	ChoosableDataType_ExactDate ChoosableDataType = 2
)
