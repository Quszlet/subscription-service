package service

import (
	"github.com/Quszlet/subscription-service/internal/repository"
	"github.com/Quszlet/subscription-service/pkg/logger"
)

// type User interface {
// 	Create(user models.User) (int, error)
// 	Update(userId int) error
// 	Get(userId int) (models.User, error)
// 	GetAll() ([]models.User, error)
// 	Delete(userId int) error
// }

type Service struct {
	
}

func NewService(r *repository.Repository, log logger.Logger) *Service {
	// return &Service{User: NewUserService(r)}
	return &Service{}
}