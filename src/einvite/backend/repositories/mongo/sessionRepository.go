package mongo

import (
	"einvite/backend/repositories"
	"einvite/common/contracts"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type sessionRepository struct {
}

type sessionEntity struct {
	Id     bson.ObjectId      `bson:"_id"`
	Values map[string]string  `bson:"values"`
	User   *sessionUserEntity `bson:"user"`
	Expiry time.Time          `bson:"expiry"`
}

type sessionUserEntity struct {
	Id       string      `bson:"_id"`
	AuthType int         `bson:"authType"`
	AuthData interface{} `bson:"authData"`
}

func (this *sessionRepository) Save(info *contracts.SessionInfo) (string, error) {

	entity := &sessionEntity{
		Values: info.Values,
		Expiry: info.Expiry,
		User: &sessionUserEntity{
			Id:       info.User.UserId,
			AuthType: int(info.User.AuthType),
			AuthData: info.User.AuthData,
		},
	}

	if info.Id == "" {
		entity.Id = bson.NewObjectId()
		return this.insertNew(entity)
	} else {
		entity.Id = bson.ObjectIdHex(info.Id)
		return this.updateExisting(entity)
	}

}

func (this *sessionRepository) insertNew(entity *sessionEntity) (string, error) {
	var sessionId string

	err := withSessions(func(sessions *mgo.Collection) error {

		_err := sessions.Insert(entity)
		if _err == nil {
			sessionId = entity.Id.Hex()
		}

		return _err
	})

	return sessionId, err
}

func (this *sessionRepository) updateExisting(entity *sessionEntity) (string, error) {

	var sessionId string

	err := withSessions(func(sessions *mgo.Collection) error {

		_, _err := sessions.Upsert(bson.M{"_id": entity.Id}, entity)

		if _err == nil {
			sessionId = entity.Id.Hex()
		}

		return _err
	})

	return sessionId, err
}

func (this *sessionRepository) Get(sessionId string) (*contracts.SessionInfo, error) {

	var info *contracts.SessionInfo

	err := withSessions(func(sessions *mgo.Collection) error {

		entity := &sessionEntity{}
		_err := sessions.FindId(bson.ObjectIdHex(sessionId)).One(&entity)

		if _err == nil {
			info = toSessionDto(entity)
		}

		return _err
	})

	return info, err
}

func (this *sessionRepository) SetExpiry(sessionId string, expiry time.Time) error {

	return withSessions(
		func(sessions *mgo.Collection) error {
			return sessions.UpdateId(bson.ObjectIdHex(sessionId), bson.M{"$set": bson.M{"expiry": expiry}})
		})
}

func (this *sessionRepository) Remove(sessionId string) error {

	return withSessions(
		func(sessions *mgo.Collection) error {
			return sessions.RemoveId(bson.ObjectIdHex(sessionId))
		})
}

func NewSessionRepository() repositories.SessionRepository {

	return &sessionRepository{}
}
