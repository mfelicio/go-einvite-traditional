package mongo

import (
	"einvite/backend/repositories"
	"einvite/common/contracts"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

const (
	EVENT_PARTICIPANT_ID = "event_%s.participant"
	EVENT_ACTIVITY_ID    = "event_%s.activity"
	EVENT_CHOOSABLE_ID   = "event_%s.choosable"
)

type eventRepository struct {
}

func (this *eventRepository) Create(eventDto *contracts.Event) (*contracts.Event, error) {

	//no need to increment db counters for each participant/activity/choosable at Create time
	//we can just set them to the last value
	//participants are created with PendingAccept status
	//participant.user is nil. it will be set when the user accepts the invite

	//TODO: missing bson attributes in data from Choosables

	eventCreatorId := 0

	var entity = &dbEvent{}

	entity.Id = bson.NewObjectId()
	entity.Name = eventDto.Name
	entity.Version = 1
	entity.Creator = &dbUser{Email: eventDto.Creator.Email, Name: eventDto.Creator.Name}

	//participants
	entity.Participants = make([]*dbParticipant, len(eventDto.Participants))

	for i, participantDto := range eventDto.Participants {
		entity.Participants[i] = &dbParticipant{
			Id:     i + 1,
			Role:   participantDto.Role,
			Status: contracts.ParticipantStatus_PendingAccept,
			Email:  participantDto.Email,
			User:   nil,
		}
	}

	//activities
	entity.Activities = make([]*dbActivity, len(eventDto.Activities))
	entity.MainActivity = 1 //TODO: should come from the activity dto which activity is the main one

	nextChoosableId := 0 //shared for all activities within the event

	getNextChoosableIdFunc := func() int {
		nextChoosableId++
		return nextChoosableId
	}

	for i, activityDto := range eventDto.Activities {

		activity := &dbActivity{
			Id:     i + 1,
			Type:   activityDto.Type,
			Owner:  eventCreatorId,
			Whats:  toDbChoosables(activityDto.Whats, getNextChoosableIdFunc),
			Whens:  toDbChoosables(activityDto.Whens, getNextChoosableIdFunc),
			Wheres: toDbChoosables(activityDto.Wheres, getNextChoosableIdFunc),
		}

		if len(activity.Whats) > 0 {
			activity.What = activity.Whats[0].Id
		}

		if len(activity.Whens) > 0 {
			activity.When = activity.Whens[0].Id
		}

		if len(activity.Wheres) > 0 {
			activity.Where = activity.Wheres[0].Id
		}

		entity.Activities[i] = activity
	}

	//persist

	var event *contracts.Event

	err := withEvents(func(events *mgo.Collection) error {

		_err := setIdsForNewEvent(events, entity.Id.String(), len(entity.Participants), len(entity.Activities), nextChoosableId)

		if _err == nil {

			_err = events.Insert(entity)

			if _err == nil {

				read := dbEvent{}
				events.FindId(entity.Id).One(&read)

				event = toEventDto(&read)
			}
		}

		return _err
	})

	return event, err
}

func NewEventRepository() repositories.EventRepository {

	return &eventRepository{}
}
