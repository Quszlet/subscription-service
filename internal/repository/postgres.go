package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	usersTable = "users"
)

type ConfigPostgres struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewConfigPostgres(host string, port string, username string, password string, DBname string, sslmode string) *ConfigPostgres {
	return &ConfigPostgres{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		DBName:   DBname,
		SSLMode:  sslmode,
	}
}

func NewPostgresDB(cfgpg *ConfigPostgres) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfgpg.Host, cfgpg.Port, cfgpg.Username, cfgpg.DBName, cfgpg.Password, cfgpg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
