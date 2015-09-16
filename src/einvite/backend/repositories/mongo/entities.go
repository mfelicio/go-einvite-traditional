package mongo

import (
	"einvite/common/contracts"
	"einvite/framework"
	"labix.org/v2/mgo/bson"
	"time"
)

/*
each event will have:

(event.id).participantId ~ ev123abc.participantId
(event.id).activityId	 ~ ev123abc.activityId
(event.id).choosableId	 ~ ev123abc.choosableId
(event.id).version		 ~ ev123abc.version

*/

type dbCounter struct {
	Name    string `bson:"_id"`
	Counter int    `bson:"counter"`
}

type dbParticipant struct {
	Id     int                         `bson:"id"`
	Email  string                      `bson:"email"`
	User   *dbUser                     `bson:"user"`
	Role   contracts.ParticipantRole   `bson:"role"`
	Status contracts.ParticipantStatus `bson:"status"` //pending-accept , accepted
}

type dbActivity struct {
	Id         int                    `bson:"id"`
	Type       contracts.ActivityType `bson:"type"`
	Owner      int                    `bson:"owner"`
	Interested []int                  `bson:"interested"`
	Wheres     []*dbChoosable         `bson:"wheres"`
	Whats      []*dbChoosable         `bson:"whats"`
	Whens      []*dbChoosable         `bson:"whens"`
	Where      int                    `bson:"where"`
	What       int                    `bson:"what"`
	When       int                    `bson:"when"`
}

type dbChoosable struct {
	Id       int                     `bson:"id"`
	Type     contracts.ChoosableType `bson:"type"`
	Owner    int                     `bson:"owner"`
	Voters   []int                   `bson:"voters"`
	DataType contracts.ChoosableDataType
	Data     interface{} `bson:"data"`
}

type dbEvent struct {
	Id           bson.ObjectId    `bson:"_id"`
	Name         string           `bson:"name"`
	Creator      *dbUser          `bson:"creator"`
	Participants []*dbParticipant `bson:"participants"`
	Activities   []*dbActivity    `bson:"activities"`
	MainActivity int              `bson:"mainActivity"`
	Version      int              `bson:"version"`
}

type dbUser struct {
	Email string      `bson:"_id"`
	Name  string      `bson:"name"`
	Auth  *dbUserAuth `bson:"auth"`
}

type dbUserAuth struct {
	Facebook *dbUserAuthCredentials `bson:"facebook"`
	Google   *dbUserAuthCredentials `bson:"google"`
	Twitter  *dbUserAuthCredentials `bson:"twitter"`
}

type dbUserAuthCredentials struct {
	Type         framework.AuthType `bson:"type"`
	AccessToken  string             `bson:"accessToken"`
	RefreshToken string             `bson:"refreshToken"`
	Expiry       time.Time          `bson:"expiry"`
}
