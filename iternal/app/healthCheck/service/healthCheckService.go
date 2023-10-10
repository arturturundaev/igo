package service

import (
	"igo/iternal/app/common/db"
	"igo/iternal/app/healthCheck/dto"
)

type HealthCheckService struct {
	repository db.IFaceHealthCheckRepository
}

func NewHealthCheckService(repository db.IFaceHealthCheckRepository) HealthCheckService {
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
