package handler

import (
	"github.com/gin-gonic/gin"
	"igo/iternal/app/healthCheck/handler"
)

type Handler struct {
	HealthCheckHandler handler.HealthCheckHandler
}

func NewHandler(HealthCheckHandler handler.HealthCheckHandler) Handler {
	return Handler{HealthCheckHandler: HealthCheckHandler}
}

func (handler Handler) Init() *gin.Engine {
	router := gin.New()
	router.GET("health/check", handler.HealthCheckHandler.GetResponse)

	return router
}
