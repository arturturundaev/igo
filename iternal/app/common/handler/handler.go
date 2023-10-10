package handler

import (
	"github.com/gin-gonic/gin"
	"igo/iternal/app/healthCheck/handler"
	handler2 "igo/iternal/app/user/handler"
)

type Handler struct {
	HealthCheckHandler handler.HealthCheckHandler
	userListHandler    handler2.UserListHandler
}

func NewHandler(HealthCheckHandler handler.HealthCheckHandler, userListHandler handler2.UserListHandler) Handler {
	return Handler{HealthCheckHandler: HealthCheckHandler, userListHandler: userListHandler}
}

func (handler Handler) Init() *gin.Engine {
	router := gin.New()
	router.GET("health/check", handler.HealthCheckHandler.GetResponse)
	router.GET("user/list", handler.userListHandler.GetAllUsers)

	return router
}
