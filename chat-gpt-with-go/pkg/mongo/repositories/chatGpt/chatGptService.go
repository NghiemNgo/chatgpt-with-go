package chatGpt

import (
	"go.tienngay/pkg/mongo/entities"
)

//Service is an interface from which our api module can access our repository of all our models
type ChatGptService interface {
	Insert(chat entities.ChatGpt) (string)
}

type chatGptService struct {
	repository ChatGptRepository
}

//NewService is used to create a single instance of the service
func NewService(r ChatGptRepository) ChatGptService {
	return &chatGptService{
		repository: r,
	}
}

//Insert is a service layer that helps fetch a transaction
func (s *chatGptService) Insert(chat entities.ChatGpt) (string) {
	return s.repository.Insert(chat)
}
