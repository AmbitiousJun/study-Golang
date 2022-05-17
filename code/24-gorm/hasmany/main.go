package main

import (
	"code/24-gorm/dao"
	"code/24-gorm/model"
	"fmt"
)

func main() {
	// 自动创建表
	dao.DB.AutoMigrate(&model.User{}, &model.Computer{})
	// 插入一个用户，同时插入它拥有的电脑
	// computers := []model.Computer{
	// 	{Name: "MacBook Pro"},
	// 	{Name: "MateBook D14"},
	// 	{Name: "MacBookAir"},
	// }
	// u := model.User{Username: "君君", Password: "111111", Computers: computers}
	// res := dao.DB.Create(&u)
	// fmt.Printf("res.RowsAffected: %v\n", res.RowsAffected)

	// 查询用户，并封装电脑数据
	u1 := model.User{}
	dao.DB.Where("username = ?", "君君").First(&u1)
	dao.DB.Model(&u1).Association("Computers").Find(&u1.Computers)
	fmt.Printf("u1: %v\n", u1)
}