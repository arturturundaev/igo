package handler

import (
	"github.com/gin-gonic/gin"
	"igo/iternal/app/user/dto"
	"igo/iternal/app/user/service"
	"net/http"
)

type UserListHandler struct {
	service service.UserService
}

func NewUserListHandler(service service.UserService) UserListHandler {
	return UserListHandler{service: service}
}

func (handler UserListHandler) GetAllUsers(context *gin.Context) {
	users := handler.service.GetAllUsers()
	var response dto.UserShortResponse
	response.SetData(users)

	context.JSON(http.StatusOK, response)
}
