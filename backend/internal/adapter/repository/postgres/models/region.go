package models

import (
	"backend/internal/core/domain"
	"time"

	"gorm.io/gorm"
)

type Region struct {
	gorm.Model
	Name         string
	Code         string
	IsActive     bool
	NextOpenTime time.Time
}

func (r Region) ToDomain() domain.Region {
	return domain.Region{
		Name:         r.Name,
		Code:         r.Code,
		IsActive:     r.IsActive,
		NextOpenTime: r.NextOpenTime,
		ID:           r.ID,
	}
}

func AsRegion(region domain.Region) Region {
	return Region{
		Name:         region.Name,
		Code:         region.Code,
		IsActive:     region.IsActive,
		NextOpenTime: region.NextOpenTime,
	}
}
