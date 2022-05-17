// 图书-类型实体，多对多关系
package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name string `gorm:"index:,unique"`
	Price float64
	Types []BookType `gorm:"many2many:book_type_rels"`
}

type BookType struct {
	gorm.Model
	Name string `gorm:"index:,unique"`
	Books []Book `gorm:"many2many:book_type_rels"`
}