package mail

import (
	"go.tienngay/pkg/mysql/entities"
)

//Service is an interface from which our api module can access our repository of all our models
type Service interface {
	GetWaitingMails() (*[]entities.Mail)
	UpdateStatusSending(ID int) (bool)
	UpdateStatusSuccess(mail entities.Mail) (bool)
	UpdateStatusErrors(mail entities.Mail) (bool)
}

type service struct {
	repository Repository
}

//NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetWaitingMails() (*[]entities.Mail) {
	return s.repository.GetWaitingMails()
}

func (s *service) UpdateStatusSending(ID int) (bool) {
	return s.repository.UpdateStatusSending(ID)
}

func (s *service) UpdateStatusSuccess(mail entities.Mail) (bool) {
	return s.repository.UpdateStatusSuccess(mail)
}

func (s *service) UpdateStatusErrors(mail entities.Mail) (bool) {
	return s.repository.UpdateStatusErrors(mail)
}