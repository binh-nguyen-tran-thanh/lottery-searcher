package port

import "backend/internal/core/domain"

type LotteryRepository interface {
	SyncResult([]domain.Result) ([]domain.Result, error)
	SaveOpenNumb([]domain.OpenNum) error
	DeleteResult(domain.Result) error
}
