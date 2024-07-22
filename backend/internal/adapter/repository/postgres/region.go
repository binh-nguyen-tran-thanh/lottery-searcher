package postgres

import (
	"backend/internal/adapter/repository/postgres/models"
	"backend/internal/core/domain"
	"backend/internal/core/port"
	"backend/internal/core/util/exception"

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

func (r regionRepo) GetRegions() (results []domain.Region, err error) {
	var regions []models.Region
	r.db.Debug()

	result := r.db.Find(&regions)

	if result.Error != nil {
		return nil, exception.Into(result.Error)
	}

	for _, region := range regions {
		results = append(results, region.ToDomain())
	}

	return results, nil
}
