package postgres

import (
	"backend/internal/adapter/repository/postgres/models"
	"backend/internal/core/domain"
	"backend/internal/core/port"
	"backend/internal/core/util"
	"backend/internal/core/util/exception"
	"time"

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

	result := r.db.Model(&models.Region{IsActive: true}).Find(&regions)

	if result.Error != nil {
		return nil, exception.Into(result.Error)
	}

	for _, region := range regions {
		results = append(results, region.ToDomain())
	}

	return results, nil
}

func (r regionRepo) GetRegionHasTurnToday() (results []domain.Region, err error) {
	var regions []models.Region

	result := r.db.Debug().Model(&models.Region{IsActive: true}).Where("next_open_time <= ?", time.Now().Local()).Find(&regions)

	if result.Error != nil {
		return nil, exception.Into(result.Error)
	}

	for _, region := range regions {
		results = append(results, region.ToDomain())
	}

	return results, nil
}

func (r regionRepo) UpdateRegionOpenTime(id uint, openTime string) (updatedCount int, err error) {
	nextOpenTime, err := util.ParseToFormattedDate(openTime)

	if err != nil {
		return 0, exception.Into(err)
	}

	result := r.db.Model(&models.Region{}).Where("id = ?", id).Update("next_open_time", nextOpenTime)

	if result.Error != nil {
		return 0, exception.Into(result.Error)
	}

	return int(result.RowsAffected), nil
}
