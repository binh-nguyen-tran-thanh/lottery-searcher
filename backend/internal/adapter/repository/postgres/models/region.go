package models

import (
	"backend/internal/core/domain"

	"gorm.io/gorm"
)

type Region struct {
	gorm.Model
	Name     string
	Code     string
	IsActive bool
}

func (r Region) ToDomain() domain.Region {
	return domain.Region{
		Name:     r.Name,
		Code:     r.Code,
		IsActive: r.IsActive,
	}
}

func AsRegion(region domain.Region) Region {
	return Region{
		Name:     region.Name,
		Code:     region.Code,
		IsActive: region.IsActive,
	}
}
