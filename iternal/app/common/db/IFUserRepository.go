package db

import "igo/iternal/app/user/model"

type IFUserRepository interface {
	GetAllUsers() []model.User
}
