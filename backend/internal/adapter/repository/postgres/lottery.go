package postgres

import (
	"backend/internal/adapter/repository/postgres/models"
	"backend/internal/core/domain"
	"backend/internal/core/port"
	"backend/internal/core/util/exception"

	"gorm.io/gorm"
)

type lotteryRepo struct {
	db *gorm.DB
}

func NewLotteryRepo(db *gorm.DB) port.LotteryRepository {
	return lotteryRepo{
		db: db,
	}
}

func (l lotteryRepo) SyncResult(result []domain.Result) (int, error) {
	if len(result) <= 0 {
		return 0, exception.New(exception.TypeInternal, "Empty result", nil)
	}

	var resultToSave []models.Result

	for idx := range result {
		resultToSave = append(resultToSave, models.AsResult(result[idx]))
	}

	res := l.db.Create(&resultToSave)

	if err := res.Error; err != nil {
		return 0, exception.New(exception.TypeInternal, "Fail to save result", err)
	}

	return int(res.RowsAffected), nil
}
