package model

import "gorm.io/gorm"

type Pet struct {
	gorm.Model
	Name string
	Age int
	Type string
}