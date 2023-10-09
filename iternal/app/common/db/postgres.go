package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	DbName   string
	Login    string
	Password string
}

// NewPostgres Get connection to Postgres
func NewPostgres(config Config) (*sqlx.DB, error) {
	database, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=%s",
		config.Host, config.Port, config.Login, config.Password, config.DbName, "disable"))

	if err != nil {
		return nil, err
	}

	if err = database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}
