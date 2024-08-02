package port

import "backend/internal/core/domain"

type LotteryRepository interface {
	SyncResult([]domain.Result) (int, error)
}
