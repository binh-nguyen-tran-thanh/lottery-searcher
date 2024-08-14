package service

import (
	"backend/internal/core/domain"
	"backend/internal/core/port"
)

type resultService struct {
	property serviceProperty
}

func NewResultService(property serviceProperty) port.ResultService {
	return resultService{
		property: property,
	}
}

func (r resultService) GetTodayResult() (returnValues []domain.Result, err error) {

	returnValues, err = r.property.repo.Lottery().FindTodayResult()

	if err != nil {
		return nil, err
	}

	return
}

func (r resultService) FilterResult(payload port.FilterResultPayload) (returnValues []domain.Result, err error) {
	returnValues, err = r.property.repo.Lottery().FilterResult(payload)

	if err != nil {
		return nil, err
	}

	return
}

func (r resultService) FilterOpenNumb(payload port.FilterOpenNumbPayload) (returnValues []domain.OpenNum, err error) {
	returnValues, err = r.property.repo.Lottery().FilterOpenNumbs(payload)

	if err != nil {
		return nil, err
	}

	return
}
