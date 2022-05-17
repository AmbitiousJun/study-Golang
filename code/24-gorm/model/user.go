package model

import "gorm.io/gorm"

// 用户实体
type User struct {
	gorm.Model
	Username string
	Password string
	// Has One
	CreditCard CreditCard
	// Has Many
	Computers []Computer
}