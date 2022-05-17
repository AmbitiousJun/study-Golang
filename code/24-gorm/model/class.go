package model

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name string
}