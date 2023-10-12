package db

import "igo/iternal/app/user/model"

type IFUserRepository interface {
	GetAllUsers(page int, limit int) []model.User
}
