package postgres

import (
	"backend/internal/core/port"
	"fmt"

	"gorm.io/gorm"
)

type regionRepo struct {
	db *gorm.DB
}

func NewRegionRepo(db *gorm.DB) port.RegionRepository {
	return regionRepo{
		db: db,
	}
}

func (r regionRepo) GetRegions() {
	fmt.Print("Place holder")
}
