package service

import (
	"igo/iternal/app/common/db"
	"igo/iternal/app/user/dto"
)

type UserService struct {
	repository db.IFUserRepository
}

func NewUserService(repository db.IFUserRepository) UserService {
	return UserService{
		repository: repository,
	}
}

func (service UserService) GetAllUsers(page int, limit int) []dto.UserShort {
	users := service.repository.GetAllUsers(page, limit)
	var usersDto []dto.UserShort
	for _, element := range users {
		usersDto = append(usersDto, dto.UserShort{First_name: element.First_name, Second_name: element.Second_name, Middle_name: element.Middle_name})
	}

	return usersDto
}
