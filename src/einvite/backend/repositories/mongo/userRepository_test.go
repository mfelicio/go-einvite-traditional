package mongo

import (
	"einvite/common/contracts"
	"fmt"
	check "gopkg.in/check.v1"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func (s *MongoTests) Test_CreateUser(c *check.C) {

	repository := NewUserRepository()

	user, err := repository.Save(&contracts.User{"user1@email.com", "user1"}, nil)

	if err != nil {

		c.Error(err.Error())
	} else {

		if user.Name != "user1" {
			c.Errorf("User name was %s and should be %s", user.Name, "user1")
		}

		if user.Email != "user1@email.com" {
			c.Errorf("User email was %s and should be %s", user.Email, "user1@email.com")
		}
	}

}

func (s *MongoTests) Test_SaveUsersWithDuplicateEmailMustNotFail(c *check.C) {

	repository := NewUserRepository()

	repository.Save(&contracts.User{"user1@email.com", "user1"}, nil)

	_, err2 := repository.Save(&contracts.User{"user1@email.com", "user1"}, nil)

	if err2 != nil {
		c.Error("Should not have failed because the Save semanthics is CreateOrUpdate")
		c.Log(err2)

	}

}

func (s *MongoTests) Test_CreateUsersWithDifferentEmailsMustNotFail(c *check.C) {

	repository := NewUserRepository()

	errors := make([]error, 4)

	_, errors[0] = repository.Save(&contracts.User{"user1@email.com", "user1"}, nil)
	_, errors[1] = repository.Save(&contracts.User{"user2@email.com", "user2"}, nil)
	_, errors[2] = repository.Save(&contracts.User{"user3@email.com", "user3"}, nil)
	_, errors[3] = repository.Save(&contracts.User{"user4@email.com", "user4"}, nil)

	for i, err := range errors {
		if err != nil {

			c.Errorf("user%d@email.com error. msg: %s", i, err.Error())
			break
		}
	}

}

func (s *MongoTests) Test_UsersHighLoad(c *check.C) {

	fmt.Println("MAXProcs", runtime.GOMAXPROCS(runtime.NumCPU()))

	num := 1000

	repository := NewUserRepository()

	randomizer := rand.NewSource(time.Now().UnixNano())

	wg := &sync.WaitGroup{}
	wg.Add(num)

	insertFunc := func(idx int) {
		defer wg.Done()
		id := randomizer.Int63() + int64(idx)

		user := &contracts.User{
			Email: fmt.Sprintf("user%d@test.com", id),
			Name:  fmt.Sprintf("User %d", id),
		}

		_, err := repository.Save(user, nil)

		if err != nil {
			fmt.Println(fmt.Sprintf("%d failed %s", idx, err.Error()))

		}
	}

	for i := 0; i < num; i++ {
		go insertFunc(i)
		//insertFunc(i, channel)
	}

	wg.Wait()

	inserted, _ := repository.Count()

	failed := num - inserted

	if failed > 0 {
		c.Errorf("Failed %d inserts", failed)
	} else {
		fmt.Println(fmt.Sprintf("Inserted %d", inserted))
	}

}
