package mongo

import (
	"fmt"
	check "gopkg.in/check.v1"
	"labix.org/v2/mgo"
	"runtime"
	"sync"
	"time"
)

func (s *MongoTests) Test_CopyCanBeUsedConcurrently(c *check.C) {

	runtime.GOMAXPROCS(runtime.NumCPU())

	masterSession, mErr := mgo.DialWithInfo(&mgo.DialInfo{
		Database: "test",
		Username: "",
		Password: "",
		Addrs:    []string{fmt.Sprintf("%s:%s", _cfg.Host, _cfg.Port)},
		Timeout:  10000 * time.Second,
	})

	if mErr != nil {
		c.Error(mErr)
		return
	}

	defer masterSession.Close()

	num := 1000
	wg := &sync.WaitGroup{}
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {

			defer wg.Done()

			s := masterSession.Copy()
			defer s.Close()

			err := s.Ping()
			if err != nil {
				c.Error(err)
			}

		}()
	}

	wg.Wait()
}
