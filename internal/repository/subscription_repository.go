package repository

import (
	"github.com/Quszlet/subscription-service/internal/models"
	"github.com/Quszlet/subscription-service/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type SubscriptionPostgres struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewSubscriptionPostgres(db *sqlx.DB, log logger.Logger) *SubscriptionPostgres {
	return &SubscriptionPostgres{db: db, log: log}
}

func (subp *SubscriptionPostgres) Create(sub models.Subscription) (int, error) {
	return 0, nil
}

func (subp *SubscriptionPostgres) Update(subId uint) error {
	return nil
}

func (subp *SubscriptionPostgres) Get(subId uint) (models.Subscription, error) {
	return models.Subscription{}, nil
}

func (subp *SubscriptionPostgres) GetAllByUser(userId uint) ([]models.Subscription, error) {
	return []models.Subscription{}, nil
}

func (subp *SubscriptionPostgres) GetByUserAndStatus(userId uint, status models.Status) ([]models.Subscription, error) {
	return []models.Subscription{}, nil
}

func (subp *SubscriptionPostgres) GetByUserAndStatusAndType(userId uint, status models.Status, typeSub models.Type) ([]models.Subscription, error) {
	return []models.Subscription{}, nil
}

func (subp *SubscriptionPostgres) 	Delete(subId uint) error {
	return nil
}
