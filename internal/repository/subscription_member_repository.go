package repository

import (
	"github.com/Quszlet/subscription-service/internal/models"
	"github.com/Quszlet/subscription-service/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type SubscriptionMemberPostgres struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewSubscriptionMemberPostgres(db *sqlx.DB, log logger.Logger) *SubscriptionMemberPostgres {
	return &SubscriptionMemberPostgres{db: db, log: log}
}

func (submemp *SubscriptionMemberPostgres) Create(subMem models.SubscriptionMember) (int, error) {
	return 0, nil
}

func (submemp *SubscriptionMemberPostgres) Update(subMemId uint) error {
	return nil
}

func (submemp *SubscriptionMemberPostgres) Get(subMemId uint) (models.SubscriptionMember, error) {
	return models.SubscriptionMember{}, nil
}

func (submemp *SubscriptionMemberPostgres) GetByOwnerId(userId uint) ([]models.SubscriptionMember, error) {
	return []models.SubscriptionMember{}, nil
}

func (submemp *SubscriptionMemberPostgres) GetBySubscriptionId(subId uint) ([]models.SubscriptionMember, error) {
	return []models.SubscriptionMember{}, nil
}

func (submemp *SubscriptionMemberPostgres) Delete(userId uint) error {
	return nil
}