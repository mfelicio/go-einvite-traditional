package mongo

import (
	"einvite/backend/repositories"
	"einvite/common/contracts"
	"einvite/framework"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type userRepository struct {
}

func (this *userRepository) Save(userDto *contracts.User, credentialsDto *contracts.UserAuthCredentials) (*contracts.User, error) {

	var user *contracts.User

	//must assign err because withUsers returns a different error
	err := withUsers(func(users *mgo.Collection) error {

		set := bson.M{
			"name": userDto.Name,
		}

		if credentialsDto != nil {
			var authtype string

			switch credentialsDto.Type {
			case framework.AuthType_Google:
				authtype = "google"
			case framework.AuthType_Facebook:
				authtype = "facebook"
			case framework.AuthType_Twitter:
				authtype = "twitter"
			}

			authfieldprefix := fmt.Sprintf("auth.%s", authtype)
			set[fmt.Sprintf("%s.type", authfieldprefix)] = credentialsDto.Type
			set[fmt.Sprintf("%s.accessToken", authfieldprefix)] = credentialsDto.AccessToken
			set[fmt.Sprintf("%s.expiry", authfieldprefix)] = credentialsDto.Expiry

			// only store refreshtoken if it isn't nil because it may not be
			// filled in each oauth login
			if credentialsDto.RefreshToken != "" {
				set[fmt.Sprintf("%s.refreshToken", authfieldprefix)] = credentialsDto.RefreshToken
			}

		}

		upsert := mgo.Change{
			Update:    bson.M{"$set": set},
			ReturnNew: true,
			Upsert:    true,
		}

		var entity *dbUser
		_, _err := users.FindId(userDto.Email).Apply(upsert, &entity)

		if _err == nil {

			user = toUserDto(entity)
		}

		return _err
	})

	return user, err
}

func (this *userRepository) Get(id string) (*contracts.User, error) {

	var user *contracts.User

	//must assign err because withUsers returns a different error
	err := withUsers(func(users *mgo.Collection) error {

		entity := &dbUser{}
		_err := users.FindId(id).One(&entity)

		if _err == nil {
			user = toUserDto(entity)
		}

		return _err
	})

	return user, err
}

func (this *userRepository) List() ([]*contracts.User, error) {

	var dtos []*contracts.User

	err := withUsers(func(users *mgo.Collection) error {

		var entities []*dbUser
		_err := users.Find(bson.M{}).All(&entities)
		if _err == nil {
			dtos = toUserDtos(entities)
		}
		return _err
	})

	return dtos, err
}

func (this *userRepository) Count() (int, error) {
	var counter int

	err := withUsers(func(users *mgo.Collection) error {
		c, _err := users.Count()

		if _err == nil {
			counter = c
		}

		return _err
	})

	return counter, err
}

func NewUserRepository() repositories.UserRepository {

	return &userRepository{}
}
