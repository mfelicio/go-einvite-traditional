package services

import (
	"einvite/backend/repositories"
	"einvite/common/contracts"
	"einvite/common/services"
)

type userService struct {
	userRepository repositories.UserRepository
}

func (this userService) List() ([]*contracts.User, error) {

	return this.userRepository.List()
}

func (this userService) Create(userDto *contracts.User) (*contracts.User, error) {

	user, err := this.userRepository.Save(userDto, nil)

	return user, err
}

func (this userService) SaveWithCredentials(userDto *contracts.User, credentialsDto *contracts.UserAuthCredentials) (*contracts.User, error) {

	user, err := this.userRepository.Save(userDto, credentialsDto)

	return user, err
}

func (this userService) Get(id string) (*contracts.User, error) {

	user, err := this.userRepository.Get(id)

	return user, err
}

func NewUserService(userRepository repositories.UserRepository) services.UserService {

	return &userService{userRepository}
}
