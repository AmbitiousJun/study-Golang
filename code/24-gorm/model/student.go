package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string
	Class Class
	ClassID uint
}