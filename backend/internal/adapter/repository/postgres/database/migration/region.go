package migration

import (
	"backend/internal/adapter/repository/postgres/models"
	"encoding/json"
	"os"
	"path/filepath"
)

type Regions = []*models.Region

func (m *Migration) SeedRegions() int {
	defer m.wg.Done()

	var count int64

	regionModel := m.db.Model(&models.Region{})

	regionModel.Count(&count)

	if count > 0 {
		m.logger.Info().Msg("Skip seeding region because of data have already existed")
		return int(count)
	}

	m.logger.Info().Msg("Reading Region Config file")

	wd, _ := os.Getwd()

	filePath := filepath.Join(wd, "/internal/adapter/repository/postgres/database/migration/seeders/regions.json")

	jsonContent, err := os.ReadFile(filePath)

	if err != nil {
		m.logger.Fatal().Msgf("Fail to read seeding file. Reason: %s", err.Error())
		return 0
	}

	var regions Regions

	if err := json.Unmarshal(jsonContent, &regions); err != nil {
		m.logger.Fatal().Msgf("Fail to parse seed file to json. Reason: %s", err.Error())
		return 0
	}

	for idx := range regions {
		regions[idx].IsActive = true
	}

	result := regionModel.Create(regions)

	if result.Error != nil {
		m.logger.Fatal().Msgf("Fail to insert to database. Reason: %s", result.Error.Error())
		return 0
	}

	m.logger.Info().Msgf("Seeded %d regions to database", result.RowsAffected)

	return int(result.RowsAffected)
}
