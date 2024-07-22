package repository

import (
	"backend/internal/adapter/repository/postgres"
	"backend/internal/core/port"
	"backend/internal/core/util"
)

func NewRepository(config util.Config, logger port.Logger) (port.Repository, error) {
	return postgres.NewPostgresRepo(config, logger)
}
