package mongo

import (
	"einvite/framework"
	"fmt"
	"labix.org/v2/mgo"
	"sync"
	"time"
)

type mongoConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string

	ConnectionString string
}

type dbSession struct {
	session *mgo.Session
	DB      *mgo.Database
}

func (this *dbSession) Close() {
	this.session.Close()
}

//master session, created only once (recommended by mgo)
//More info here: http://denis.papathanasiou.org/2012/10/14/go-golang-and-mongodb-using-mgo/
var _master *mgo.Session
var _cfg *mongoConfig //fields are populated the first time the configuration is read
var _lock = &sync.Mutex{}

func getSession() (*mgo.Session, error) {

	_lock.Lock()
	defer _lock.Unlock()

	if _master == nil {
		fmt.Println("Create master dbSession")
		var err error
		if _master, err = _createSession(); err != nil {
			return nil, err
		}

		fmt.Println("Ensuring db indexes")
		_ensureDbIndexes(_master)
	}

	//return _master.Clone(), nil
	return _master.Copy(), nil
	//return _createSession()
	//return _master, nil
	//return _master.Copy(), nil
}

func getDbSession() (*dbSession, error) {

	var err error
	var session *mgo.Session

	if session, err = getSession(); err == nil {

		//validate connection
		//if err = session.Ping(); err == nil {
		return &dbSession{session, session.DB(_cfg.Database)}, nil
		//}
	}

	return nil, err
}

func _ensureDbIndexes(session *mgo.Session) {

	sessions := session.DB(_cfg.Database).C(SESSIONS_COLLECTION)
	expireSessionsIndex := mgo.Index{}
	//TODO: should read from the sessions config
	//or be invoked from the bootstrapper
	expireSessionsIndex.ExpireAfter = 1 * time.Hour
	expireSessionsIndex.Key = []string{"expiry"}
	sessions.EnsureIndex(expireSessionsIndex)
}

//private to this file
func _createSession() (*mgo.Session, error) {
	
	if _cfg == nil {
		_cfg = &mongoConfig{}
		framework.Config.ReadInto("mongo", &_cfg)
	}

	var timeout = 10 * time.Second

	if _cfg.ConnectionString != "" {

		return mgo.DialWithTimeout(_cfg.ConnectionString, timeout)
	} else {

		dialInfo := mgo.DialInfo{}
		dialInfo.Database = _cfg.Database
		dialInfo.Username = _cfg.User
		dialInfo.Password = _cfg.Password
		dialInfo.Addrs = []string{fmt.Sprintf("%s:%s", _cfg.Host, _cfg.Port)}
		dialInfo.Timeout = timeout

		fmt.Println("Logging to", dialInfo.Addrs[0])

		return mgo.DialWithInfo(&dialInfo)
	}
}
