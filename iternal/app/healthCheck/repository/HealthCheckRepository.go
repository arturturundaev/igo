package repository

import (
	"github.com/jmoiron/sqlx"
)

type HealthCheckRepository struct {
	db *sqlx.DB
}

type isActive struct {
	IsActive bool
}

func NewHealthCheckRepository(db *sqlx.DB) HealthCheckRepository {
	return HealthCheckRepository{
		db: db,
	}
}

func (repo HealthCheckRepository) Check() bool {
	var isActive []isActive
	err := repo.db.Select(&isActive, "SELECT 1 as IsActive")
	if err != nil {
		return false
	}

	return true
}
