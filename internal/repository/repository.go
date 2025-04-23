package repository

import (
	"github.com/Quszlet/subscription-service/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
}

func NewRepository(db *sqlx.DB, log logger.Logger) *Repository {
	return &Repository{}
}