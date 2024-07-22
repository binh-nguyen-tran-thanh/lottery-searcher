package service

import (
	"backend/internal/core/port"
)

type pingService struct {
	property serviceProperty
}

func NewPingService(property serviceProperty) port.IPing {
	return pingService{
		property: property,
	}
}

func (s pingService) Ping() (map[string]any, error) {
	resp := make(map[string]any)
	resp["message"] = "Pong"
	return resp, nil
}

func (s pingService) Healthy() (map[string]any, error) {
	resp := make(map[string]any)
	resp["message"] = "I'm good"
	return resp, nil
}
