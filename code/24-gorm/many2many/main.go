// many2many的使用
package main

import (
	"code/24-gorm/dao"
	"code/24-gorm/model"
	"fmt"
)

func main() {
	// 自动创建表
	dao.DB.AutoMigrate(&model.Book{}, &model.BookType{})

	// 插入类型
	// types := []model.BookType{
	// 	{Name: "数学"},
	// 	{Name: "物理"},
	// 	{Name: "计算机"},
	// 	{Name: "英语"},
	// }
	// res := dao.DB.Create(&types)
	// fmt.Printf("res.RowsAffected: %v\n", res.RowsAffected)

	// 插入图书
	// types := []model.BookType{}
	// mathType := model.BookType{Name: "数学"}
	// dao.DB.First(&mathType)
	// types = append(types, mathType)
	// book := model.Book{Name: "高等数学1", Price: 37.99, Types: types}
	// res := dao.DB.Create(&book)
	// fmt.Printf("res.RowsAffected: %v\n", res.RowsAffected)

	// 查找图书
	book := model.Book{Name: "高等数学1"}
	dao.DB.First(&book)
	dao.DB.Model(&book).Association("Types").Find(&book.Types)
	fmt.Printf("book: %v\n", book)
}
