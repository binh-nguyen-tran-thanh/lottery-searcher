package port

import "backend/internal/core/domain"

type ResultService interface {
	GetTodayResult() ([]domain.Result, error)
	FilterResult(FilterResultPayload) ([]domain.Result, error)
	FilterOpenNumb(FilterOpenNumbPayload) ([]domain.OpenNum, error)
}
