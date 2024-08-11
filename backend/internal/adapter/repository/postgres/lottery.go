package postgres

import (
	"backend/internal/adapter/repository/postgres/models"
	"backend/internal/core/domain"
	"backend/internal/core/port"
	"backend/internal/core/util/exception"

	"gorm.io/gorm"
)

type lotteryRepo struct {
	db     *gorm.DB
	logger port.Logger
}

func NewLotteryRepo(db *gorm.DB, logger port.Logger) port.LotteryRepository {
	return lotteryRepo{
		db:     db,
		logger: logger,
	}
}

func (l lotteryRepo) SyncResult(result []domain.Result) (returnValues []domain.Result, err error) {
	if len(result) <= 0 {
		return nil, exception.New(exception.TypeInternal, "Empty result", nil)
	}

	var resultToSave []models.Result

	for idx := range result {
		resultToSave = append(resultToSave, models.AsResult(result[idx]))
	}

	res := l.db.Create(&resultToSave)

	if err := res.Error; err != nil {
		return nil, exception.New(exception.TypeInternal, "Fail to save result", err)
	}

	for _, v := range resultToSave {
		returnValues = append(returnValues, v.ToDomain())
	}

	return returnValues, nil
}

func (l lotteryRepo) SaveOpenNumb(data []domain.OpenNum) error {
	if len(data) <= 0 {
		return exception.New(exception.TypeNotFound, "Empty value", nil)
	}

	var valueToSave []models.OpenNumb

	for _, v := range data {
		valueToSave = append(valueToSave, models.AsOpenNumb(v))
	}

	result := l.db.Create(valueToSave)

	if result.Error != nil {
		return exception.New(exception.TypeInternal, "Fail to save", result.Error)
	}

	return nil
}

func (l lotteryRepo) DeleteResult(result domain.Result) error {
	res := l.db.Delete(&result)

	if res.Error != nil {
		return exception.New(exception.TypeInternal, "Delete failed", res.Error)
	}

	return nil
}
