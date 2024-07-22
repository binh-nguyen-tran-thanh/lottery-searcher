package service

import (
	"backend/internal/core/port"
	"backend/internal/core/util"
)

type serviceProperty struct {
	config util.Config
	logger port.Logger
	repo   port.Repository
}

type services struct {
	property      serviceProperty
	pingService   port.IPing
	regionService port.RegionServicer
}

func NewService(config util.Config, logger port.Logger, repo port.Repository) (port.Service, error) {
	property := serviceProperty{
		config: config,
		logger: logger,
		repo:   repo,
	}

	svc := services{
		property:      property,
		pingService:   NewPingService(property),
		regionService: NewRegionService(property),
	}

	return svc, nil
}

func (s services) Ping() port.IPing {
	return s.pingService
}

func (s services) Region() port.RegionServicer {
	return s.regionService
}
