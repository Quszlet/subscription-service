package repository

import (
	"github.com/Quszlet/subscription-service/internal/models"
	"github.com/Quszlet/subscription-service/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(user models.User) (int, error)
	Update(userId int) error
	Get(userId int) (models.User, error)
	Delete(userId int) error
}

type Subscription interface {
	Create(sub models.Subscription) (int, error)
	Update(subId uint) error
	Get(subId uint) (models.Subscription, error)
	GetAllByUser(userId uint) ([]models.Subscription, error)
	GetByUserAndStatus(userId uint, status models.Status) ([]models.Subscription, error)
	GetByUserAndStatusAndType(userId uint, status models.Status, typeSub models.Type) ([]models.Subscription, error)
	Delete(subId uint) error
}

type SubscriptionMember interface {
	Create(subMem models.SubscriptionMember) (int, error)
	Update(subMemId uint) error
	Get(subMemId uint) (models.SubscriptionMember, error)
	GetByOwnerId(userId uint) ([]models.SubscriptionMember, error)
	GetBySubscriptionId(subId uint) ([]models.SubscriptionMember, error)
	Delete(userId uint) error
}

type Repository struct {
	User
	Subscription
	SubscriptionMember
}

func NewRepository(db *sqlx.DB, log logger.Logger) *Repository {
	return &Repository{
		User:               NewUserPostgres(db, log),
		Subscription:       NewSubscriptionPostgres(db, log),
		SubscriptionMember: NewSubscriptionMemberPostgres(db, log),
	}
}
