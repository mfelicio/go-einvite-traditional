package repositories

import (
	contracts "einvite/common/contracts"
)

type UserRepository interface {
	Save(userDto *contracts.User, credentialsDto *contracts.UserAuthCredentials) (*contracts.User, error)
	Get(id string) (*contracts.User, error)
	List() ([]*contracts.User, error)
	Count() (int, error)
}
