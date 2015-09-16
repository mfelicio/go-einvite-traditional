package services

import (
	"einvite/backend/repositories"
	"einvite/common/contracts"
	"einvite/common/services"
	"time"
)

type sessionService struct {
	sessionRepository repositories.SessionRepository
}

func (this *sessionService) Save(info *contracts.SessionInfo) (string, error) {
	return this.sessionRepository.Save(info)
}

func (this *sessionService) Get(sessionId string) (*contracts.SessionInfo, error) {
	return this.sessionRepository.Get(sessionId)
}

func (this *sessionService) SetExpiry(sessionId string, expiry time.Time) error {
	return this.sessionRepository.SetExpiry(sessionId, expiry)
}

func (this *sessionService) Remove(sessionId string) error {
	return this.sessionRepository.Remove(sessionId)
}

func NewSessionService(repository repositories.SessionRepository) services.SessionService {
	return &sessionService{sessionRepository: repository}
}
