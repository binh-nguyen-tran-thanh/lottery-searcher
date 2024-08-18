package port

import (
	"backend/internal/core/domain"
	"time"
)

type PagingParamPayload struct {
	Limit  uint
	Offset uint
}

type FilterResultPayload struct {
	PagingParamPayload
	FilterRangeFrom time.Time
	FilterRangeTo   time.Time
	Region          string
	ID              uint
}

type FilterOpenNumbPayload struct {
	PagingParamPayload
	FilterTye       int
	FilterRangeFrom time.Time
	FilterRangeTo   time.Time
	Region          string
	ID              uint
	ResultId        uint
	FilterMode      string
	FilterValue     string
}

type LotteryRepository interface {
	SyncResult([]domain.Result) ([]domain.Result, error)
	SaveOpenNumb([]domain.OpenNum) error
	DeleteResult(domain.Result) error
	FilterOpenNumbs(FilterOpenNumbPayload) ([]domain.OpenNum, error)
	FilterResult(FilterResultPayload) ([]domain.Result, error)
	FindTodayResult() ([]domain.Result, error)
}
