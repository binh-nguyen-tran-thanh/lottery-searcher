package domain

import "time"

type Region struct {
	ID           uint
	Name         string
	Code         string
	IsActive     bool
	NextOpenTime time.Time
}

func NewRegion(arg Region) *Region {
	return &Region{
		ID:           arg.ID,
		Name:         arg.Name,
		Code:         arg.Code,
		IsActive:     arg.IsActive,
		NextOpenTime: arg.NextOpenTime,
	}
}
