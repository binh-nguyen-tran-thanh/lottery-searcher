package migration

import (
	"backend/internal/adapter/repository/postgres/models"
	"backend/internal/core/port"
	"sync"

	"gorm.io/gorm"
)

type Migration struct {
	db     *gorm.DB
	logger port.Logger
	wg     *sync.WaitGroup
}

func New(db *gorm.DB, logger port.Logger) *Migration {
	return &Migration{
		db:     db,
		logger: logger,
		wg:     &sync.WaitGroup{},
	}
}

func (m Migration) Start() error {
	m.logger.Info().Msg("Start Migration ...")
	if err := m.db.AutoMigrate(&models.Region{}, &models.Result{}, &models.OpenNumb{}); err != nil {
		return err
	}

	m.wg.Add(1)

	go m.SeedRegions()

	m.wg.Wait()

	m.logger.Info().Msg("End Migration ...")

	return nil
}
