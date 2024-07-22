package service

import (
	"backend/internal/core/port"
	"backend/internal/core/util"
)

type serviceProperty struct {
	config util.Config
	logger port.Logger
}

type services struct {
	property    serviceProperty
	pingService port.IPing
}

func NewService(config util.Config, logger port.Logger) (port.Service, error) {
	property := serviceProperty{
		config: config,
		logger: logger,
	}

	svc := services{
		property:    property,
		pingService: NewPingService(property),
	}

	return svc, nil
}

func (s services) Ping() port.IPing {
	return s.pingService
}
