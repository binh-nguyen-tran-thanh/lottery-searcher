package service

import (
	"backend/internal/core/domain"
	"backend/internal/core/port"
)

type regionService struct {
	property serviceProperty
}

func NewRegionService(property serviceProperty) port.RegionServicer {
	return regionService{
		property: property,
	}
}

func (s regionService) GetAllRegion() ([]domain.Region, error) {
	regions, err := s.property.repo.Region().GetRegions()
	if err != nil {
		return nil, err
	}

	return regions, nil
}
