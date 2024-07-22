package handler

import (
	"backend/internal/adapter/handler/restful"
	"backend/internal/core/port"
	"backend/internal/core/util"
)

func NewServer(config util.Config, service port.Service, logger port.Logger, repo port.Repository) port.Server {
	return restful.NewServer(config, service, logger)
}
