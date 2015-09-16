package repositories

import (
	"einvite/common/contracts"
	"time"
)

type SessionRepository interface {
	Save(info *contracts.SessionInfo) (string, error)
	Get(sessionId string) (*contracts.SessionInfo, error)
	SetExpiry(sessionId string, expiry time.Time) error
	Remove(sessionId string) error
}
