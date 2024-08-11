package postgres

import (
	"backend/internal/adapter/repository/postgres/database"
	"backend/internal/core/port"
	"backend/internal/core/util"

	"gorm.io/gorm"
)

type postgresRepo struct {
	db          *gorm.DB
	logger      port.Logger
	regionRepo  port.RegionRepository
	lotteryRepo port.LotteryRepository
}

func NewPostgresRepo(config util.Config, logger port.Logger) (port.Repository, error) {
	db, err := database.New(config, logger)
	if err != nil {
		return nil, err
	}

	return create(db.Database(), logger), nil
}

func create(db *gorm.DB, logger port.Logger) port.Repository {
	return postgresRepo{
		db:          db,
		regionRepo:  NewRegionRepo(db),
		lotteryRepo: NewLotteryRepo(db, logger),
		logger:      logger,
	}
}

func (r postgresRepo) Region() port.RegionRepository {
	return r.regionRepo
}
func (r postgresRepo) Lottery() port.LotteryRepository {
	return r.lotteryRepo
}
