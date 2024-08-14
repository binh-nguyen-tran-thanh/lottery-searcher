package models

import (
	"backend/internal/core/domain"

	"gorm.io/gorm"
)

type OpenNumb struct {
	gorm.Model
	ResultID uint `gorm:"index"`
	Result   Result
	Numbs    string
	Rank     int8 `gorm:"index"`
}

func (r OpenNumb) ToDomain() domain.OpenNum {
	return domain.OpenNum{
		ID:       r.ID,
		ResultID: r.ResultID,
		Numbs:    r.Numbs,
		Rank:     r.Rank,
		Result:   r.Result.ToDomain(),
	}
}

func AsOpenNumb(r domain.OpenNum) OpenNumb {
	return OpenNumb{
		ResultID: r.ResultID,
		Numbs:    r.Numbs,
		Rank:     r.Rank,
		Result:   AsResult(r.Result),
	}
}
