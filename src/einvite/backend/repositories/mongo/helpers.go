package mongo

import (
	"einvite/framework"
	"fmt"
	mgo "labix.org/v2/mgo"
	bson "labix.org/v2/mgo/bson"
)

func getFrameworkError(err error) *framework.FrameworkError {

	//err.(type) gets the type of err and can only be used in a switch (aka type switch)
	switch err.(type) {

	case *mgo.LastError:
		return fromMongoError(err.(*mgo.LastError))
	default:
		return framework.ToError(framework.Error_Generic, err)
	}
}

func fromMongoError(err *mgo.LastError) *framework.FrameworkError {

	var errorCode framework.FrameworkErrorCode

	switch err.Code {
	case 11000:
		errorCode = framework.Error_Db_DuplicateId
	default:
		errorCode = framework.Error_Generic
	}

	return framework.ToError(errorCode, err)
}

func _getNextId(withCollection *mgo.Collection, name string) int {

	counters := withCollection.Database.C(COUNTERS_COLLECTION)

	counter := dbCounter{}
	increment := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"counter": 1}},
		ReturnNew: true,
		Upsert:    true,
	}

	counters.Find(bson.M{"_id": name}).Apply(increment, &counter)

	return counter.Counter
}

func setIdsForNewEvent(withCollection *mgo.Collection, eventId string, lastParticipantId int, lastActivityId int, lastChoosableId int) error {
	counters := withCollection.Database.C(COUNTERS_COLLECTION)

	parcipant := &dbCounter{
		Counter: lastParticipantId,
		Name:    _getEventCounterName(eventId, EVENT_PARTICIPANT_ID),
	}

	activity := &dbCounter{
		Counter: lastActivityId,
		Name:    _getEventCounterName(eventId, EVENT_ACTIVITY_ID),
	}

	choosable := &dbCounter{
		Counter: lastChoosableId,
		Name:    _getEventCounterName(eventId, EVENT_CHOOSABLE_ID),
	}

	err := counters.Insert(parcipant, activity, choosable)

	return err
}

func getNextParticipantId(withCollection *mgo.Collection, eventId string) int {

	return _getNextId(withCollection, _getEventCounterName(eventId, EVENT_PARTICIPANT_ID))
}

func getNextActivityId(withCollection *mgo.Collection, eventId string) int {

	return _getNextId(withCollection, _getEventCounterName(eventId, EVENT_ACTIVITY_ID))
}

func getNextChoosableId(withCollection *mgo.Collection, eventId string) int {

	return _getNextId(withCollection, _getEventCounterName(eventId, EVENT_CHOOSABLE_ID))
}

func _getEventCounterName(eventId string, counterFormat string) string {
	return fmt.Sprintf(counterFormat, eventId)
}
