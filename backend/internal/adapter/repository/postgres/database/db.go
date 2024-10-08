package database

import (
	"backend/internal/adapter/repository/postgres/database/migration"
	"backend/internal/core/port"
	"backend/internal/core/util"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	config util.Config
	logger port.Logger
	db     *gorm.DB
}

func New(config util.Config, logger port.Logger) (*DB, error) {
	var err error

	database := &DB{
		config: config,
		logger: logger,
	}

	if database.db, err = database.connect(); err != nil {
		return nil, err
	}

	migrator := migration.New(database.db, logger)

	if err = migrator.Start(); err != nil {
		return nil, err
	}

	return database, nil
}

func (db DB) Database() *gorm.DB {
	return db.db
}

func (db DB) connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh", db.config.Database.Host, db.config.Database.User, db.config.Database.Password, db.config.Database.Name, db.config.Database.Port)

	logLever := logger.Error

	if db.config.Environment != "Prod" {
		logLever = logger.Info
	}

	pg, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logLever,
				Colorful: true,
			},
		),
	})

	if err != nil {
		return nil, err
	}

	return pg, nil
}
