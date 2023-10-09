package service

import (
	"igo/iternal/app/healthCheck/dto"
	"igo/iternal/app/healthCheck/repository"
)

type HealthCheckService struct {
	repository repository.HealthCheckRepository
}

func NewHealthCheckService(repository repository.HealthCheckRepository) HealthCheckService {
	service := HealthCheckService{
		repository: repository,
	}

	return service
}

func (service HealthCheckService) GetAll() dto.Response {

	response := dto.Response{
		IsActiveConnectToDB: service.isIsActiveConnectionToDB(),
	}

	return response
}

func (service HealthCheckService) isIsActiveConnectionToDB() bool {
	return service.repository.Check()
}
