package services

import (
	"einvite/common/services"
)

type mailService struct {
}

func (this *mailService) SendInvite() {

}

func NewMailService() services.MailService {
	return &mailService{}
}
