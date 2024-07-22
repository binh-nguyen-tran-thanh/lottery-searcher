package models

import "gorm.io/gorm"

type Region struct {
	gorm.Model
	Name     string
	Code     string
	IsActive bool
}
