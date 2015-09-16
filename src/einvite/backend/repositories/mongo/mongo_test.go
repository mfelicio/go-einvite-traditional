package mongo

import (
	"einvite/framework"
	check "gopkg.in/check.v1"
	"os"
	"testing"
)

// gocheck wire up
func Test(t *testing.T) {
	check.TestingT(t)
}

type MongoTests struct{}

var _ = check.Suite(&MongoTests{})

func (s *MongoTests) SetUpSuite(c *check.C) {

	framework.Config.Init("../../../config_test.json")

	if _cfg == nil {
		_cfg = &mongoConfig{}
	}

	framework.Config.ReadInto("mongo", &_cfg)

	if os.Getenv("WERCKER") != "" {
		_cfg.Host = os.Getenv("WERCKER_MONGODB_HOST")
		_cfg.Port = os.Getenv("WERCKER_MONGODB_PORT")
	}
}

func (s *MongoTests) SetUpTest(c *check.C) {

}

func (s *MongoTests) TearDownTest(c *check.C) {

	var session, _ = getDbSession()
	defer session.Close()

	session.DB.C(USERS_COLLECTION).RemoveAll(nil)
	session.DB.C(COUNTERS_COLLECTION).RemoveAll(nil)
	session.DB.C(EVENTS_COLLECTION).RemoveAll(nil)
}

func (s *MongoTests) TearDownSuite(c *check.C) {

}
