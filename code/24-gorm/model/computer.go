package model

import "gorm.io/gorm"

type Computer struct {
	gorm.Model
	Name string
	UserID uint
}