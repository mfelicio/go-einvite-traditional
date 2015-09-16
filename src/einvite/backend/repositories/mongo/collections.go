package mongo

import mgo "labix.org/v2/mgo"

const (
	USERS_COLLECTION    = "users"
	EVENTS_COLLECTION   = "events"
	COUNTERS_COLLECTION = "counters"
	SESSIONS_COLLECTION = "sessions"
)

func withUsers(fn func(*mgo.Collection) error) error {
	return _withCollection(USERS_COLLECTION, fn)
}

func withEvents(fn func(*mgo.Collection) error) error {
	return _withCollection(EVENTS_COLLECTION, fn)
}

func withSessions(fn func(*mgo.Collection) error) error {
	return _withCollection(SESSIONS_COLLECTION, fn)
}

func _withCollection(name string, fn func(*mgo.Collection) error) error {

	var session *dbSession
	var err error
	if session, err = getDbSession(); err != nil {
		return getFrameworkError(err)
	}

	defer session.Close()

	collection := session.DB.C(name)

	if err = fn(collection); err != nil {
		return getFrameworkError(err)
	}
	//session.Close()
	//session.session.Refresh()
	return nil //no error
}
