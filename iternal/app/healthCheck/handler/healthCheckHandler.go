package handler

import (
	"github.com/gin-gonic/gin"
	"igo/iternal/app/healthCheck/service"
	"net/http"
)

type HealthCheckHandler struct {
	service service.HealthCheckService
}

func NewHealthCheckHandler(service service.HealthCheckService) HealthCheckHandler {
	return HealthCheckHandler{service: service}
}

func (handler HealthCheckHandler) GetResponse(context *gin.Context) {
	context.JSON(http.StatusOK, handler.service.GetAll())
}
