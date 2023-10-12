package repository

import (
	"fmt"
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

func (repository UserRepository) GetAllUsers(page int, limit int) []model.User {
	page = page - 1
	offset := page * limit
	var result []model.User

	query := fmt.Sprintf("SELECT id, first_name, second_name, middle_name FROM users ORDER BY id DESC LIMIT %d OFFSET %d", limit, offset)
	err := repository.db.Select(&result, query)

	if err != nil {
		log.Println("Fail get data from DB")
	}
	return result
}
