// golang快速入门程序
package main

import (
	"code/24-gorm/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

// 指定user表的表名
// func (User) TableName() string { return "user_tbl" }

func main() {
	dsn := "root:cyj070723@(localhost:3306)/go_db?charset=utf8mb4"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Migrate（自动创建一张与结构体对应的表）
	d.AutoMigrate(&model.Pet{})
	// Create
	// res := d.Create(&User{Username: "王五", Password: "123456"})
	// row := res.RowsAffected
	// fmt.Println("影响行数：", row)

	// 创建几条宠物数据
	pets := []model.Pet{
		{Name: "小狗", Age: 3, Type: "类型1"},
		{Name: "小猫", Age: 4, Type: "类型2"},
		{Name: "小蛇", Age: 5, Type: "类型3"},
		{Name: "小七", Age: 6, Type: "类型1"},
		{Name: "小崔", Age: 7, Type: "类型2"},
		{Name: "小招", Age: 8, Type: "类型3"},
		{Name: "小石", Age: 9, Type: "类型1"},
		{Name: "小史", Age: 10, Type: "类型2"},
		{Name: "小往", Age: 11, Type: "类型3"},
		{Name: "小巨", Age: 12, Type: "类型1"},
	}
	res := d.Create(&pets)
	fmt.Printf("res.RowsAffected: %v\n", res.RowsAffected)
}
