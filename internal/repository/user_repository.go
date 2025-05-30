package repository

import (
	"errors"
	"fmt"

	"github.com/Quszlet/subscription-service/internal/models"
	"github.com/Quszlet/subscription-service/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewUserPostgres(db *sqlx.DB, log logger.Logger) *UserPostgres {
	return &UserPostgres{db: db, log: log}
}

func (up *UserPostgres) Create(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, email, password) values ($1, $2, $3) RETURNING id", usersTable)

	row := up.db.QueryRow(query, user.Login, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (up *UserPostgres) Update(userId int) error {
	return nil
}

func (up *UserPostgres) Get(userId int) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)
	err := up.db.Get(&user, query, userId)
	return user, err
}

func (up *UserPostgres) Delete(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)
	res, err := up.db.Exec(query, userId)
	if err != nil {
		return err
	}

	affRows, err := res.RowsAffected()
	if affRows == 0 {
		return errors.New("User with this ID does not exist")
	}

	return err
}
