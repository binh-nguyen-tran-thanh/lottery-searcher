package models

import (
	"backend/internal/core/domain"

	"gorm.io/gorm"
)

type Result struct {
	gorm.Model
	TurnNum  string
	OpenNum  string
	OpenTime string
	Detail   string
	Region   string `gorm:"index"`
}

func (r Result) ToDomain() domain.Result {
	return domain.Result{
		TurnNum:  r.TurnNum,
		OpenNum:  r.OpenNum,
		OpenTime: r.OpenTime,
		Detail:   r.Detail,
		Region:   r.Region,
	}
}

func AsResult(r domain.Result) Result {
	return Result{
		TurnNum:  r.TurnNum,
		OpenNum:  r.OpenNum,
		OpenTime: r.OpenTime,
		Detail:   r.Detail,
		Region:   r.Region,
	}
}
