package mongo

import (
	"einvite/common/contracts"
	check "gopkg.in/check.v1"
	"time"
)

func (s *MongoTests) Test_CreateEvent(c *check.C) {

	event := &contracts.Event{}
	event.Name = "My testable event"
	event.Creator = &contracts.User{Email: "my@test.com", Name: "Myself"}

	event.Participants = []*contracts.Participant{
		&contracts.Participant{
			Email: "p1@test.com",
			Role:  contracts.ParticipantRole_Contributor,
		},
		&contracts.Participant{
			Email: "p2@test.com",
			Role:  contracts.ParticipantRole_Guest,
		},
		&contracts.Participant{
			Email: "p3@test.com",
			Role:  contracts.ParticipantRole_Guest,
		},
	}

	event.MainActivity = 1
	event.Activities = []*contracts.Activity{
		&contracts.Activity{
			Type: contracts.ActivityType_Food,
			Whats: []*contracts.Choosable{
				&contracts.Choosable{
					Owner:    0,
					Type:     contracts.ChoosableType_What,
					DataType: contracts.ChoosableDataType_FreeText,
					Data: &contracts.FreeText{
						Text: "Lasagna",
					},
				},
				&contracts.Choosable{
					Owner:    0,
					Type:     contracts.ChoosableType_What,
					DataType: contracts.ChoosableDataType_FreeText,
					Data: &contracts.FreeText{
						Text: "Canelonni",
					},
				},
			},
			Wheres: []*contracts.Choosable{
				&contracts.Choosable{
					Owner:    0,
					Type:     contracts.ChoosableType_Where,
					DataType: contracts.ChoosableDataType_FreeText,
					Data: &contracts.FreeText{
						Text: "Best italian restaurant in the world",
					},
				},
			},
			Whens: []*contracts.Choosable{
				&contracts.Choosable{
					Owner:    0,
					Type:     contracts.ChoosableType_When,
					DataType: contracts.ChoosableDataType_ExactDate,
					Data: &contracts.ExactDate{
						Date: time.Now().AddDate(0, 0, 1), //tomorrow
					},
				},
				&contracts.Choosable{
					Owner:    0,
					Type:     contracts.ChoosableType_When,
					DataType: contracts.ChoosableDataType_ExactDate,
					Data: &contracts.ExactDate{
						Date: time.Now().AddDate(0, 0, 2), //after tomorrow
					},
				},
			},
		},
		&contracts.Activity{
			Type: contracts.ActivityType_Movie,
			Whats: []*contracts.Choosable{
				&contracts.Choosable{
					Owner:    0,
					Type:     contracts.ChoosableType_What,
					DataType: contracts.ChoosableDataType_FreeText,
					Data: &contracts.FreeText{
						Text: "Movie 1",
					},
				},
				&contracts.Choosable{
					Owner:    0,
					Type:     contracts.ChoosableType_What,
					DataType: contracts.ChoosableDataType_FreeText,
					Data: &contracts.FreeText{
						Text: "Movie 2",
					},
				},
			},
			Wheres: []*contracts.Choosable{
				&contracts.Choosable{
					Owner:    0,
					Type:     contracts.ChoosableType_Where,
					DataType: contracts.ChoosableDataType_FreeText,
					Data: &contracts.FreeText{
						Text: "Best cinema in the world",
					},
				},
			},
			Whens: []*contracts.Choosable{
				&contracts.Choosable{
					Owner:    0,
					Type:     contracts.ChoosableType_When,
					DataType: contracts.ChoosableDataType_ExactDate,
					Data: &contracts.ExactDate{
						Date: time.Now().AddDate(0, 0, 1), //tomorrow
					},
				},
			},
		},
	}

	var repository = NewEventRepository()

	evt, err := repository.Create(event)

	if err != nil {
		c.Error(err)
	} else {
		c.Log(evt)
	}

}
