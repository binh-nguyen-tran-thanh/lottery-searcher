package migration

import (
	"backend/internal/adapter/handler/restful/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Region{}); err != nil {
		return err
	}

	return nil
}
