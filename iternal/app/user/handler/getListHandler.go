package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"igo/iternal/app/user/dto"
	"igo/iternal/app/user/service"
	"net/http"
	"strconv"
)

var validate *validator.Validate

type UserListHandler struct {
	service service.UserService
}

func NewUserListHandler(service service.UserService) UserListHandler {
	return UserListHandler{service: service}
}

func (handler UserListHandler) GetAllUsers(context *gin.Context) {
	var response dto.UserShortResponse

	page, _ := strconv.Atoi(context.Query("page"))
	limit, _ := strconv.Atoi(context.Query("limit"))
	userShortRequest := dto.UserShortRequest{Limit: limit, Page: page}
	validate = validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(userShortRequest)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}

		context.JSON(http.StatusOK, response)
		return
	}

	users := handler.service.GetAllUsers(userShortRequest.Page, userShortRequest.Limit)
	response.SetData(users)

	context.JSON(http.StatusOK, response)
}
