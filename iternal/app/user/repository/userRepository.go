package repository

import (
	"github.com/jmoiron/sqlx"
	"igo/iternal/app/user/model"
	"log"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (repository UserRepository) GetAllUsers() []model.User {
	var result []model.User

	err := repository.db.Select(&result, "SELECT id, first_name, second_name, middle_name FROM users")

	if err != nil {
		log.Println("Fail get data from DB")
	}
	return result
}
