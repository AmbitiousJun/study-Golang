package main

import (
	"code/24-gorm/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// 按照map进行插入
func createByMap() {
	rows := db.Model(&model.User{}).Create(map[string]interface{}{
		"Username": "张三", "Password": "123456789",	
	})
	fmt.Println(rows.RowsAffected)
}

// 按批次批量插入数据
func createInBatches() {
	users := []model.User{
		{Username: "张三", Password: "123456"},
		{Username: "李四", Password: "123456"},
	}
	// 按批次批量插入，每次插入1条
	rows := db.CreateInBatches(&users, 1)
	fmt.Println(rows.RowsAffected)
}

// 批量插入数据
func createMany() {
	users := []model.User{
		{Username: "张三", Password: "123456"},
		{Username: "李四", Password: "123456"},
		{Username: "王五", Password: "123456"},
	}
	rows := db.Create(&users)
	fmt.Println(rows.RowsAffected)
}

// 指定字段进行插入
func createByGivingColumns() {
	u := model.User{Username: "张三", Password: "345345"}
	res := db.Select("Username").Create(&u)
	row := res.RowsAffected
	fmt.Println("影响行数:", row)
}

// 不插入指定字段
func createAvoidingColumns() {
	u := model.User{Username: "田七", Password: "121212"}
	res := db.Omit("Password").Create(&u)
	row := res.RowsAffected
	fmt.Println("影响行数:", row)
}

func main() {
	dsn := "root:cyj070723@tcp(localhost:3306)/go_db?charset=utf8mb4"
	d, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	db = d
	db.AutoMigrate(&model.Pet{}, &model.User{})
	// 插入一个实体（在hello程序中已经演示）
	// 按照指定字段插入实体
	// createByGivingColumns()

	// 不插入指定字段
	// createAvoidingColumns()

	// 批量插入数据
	// createMany()

	// 按批量插入数据
	// createInBatches()

	// 按照map插入
	// createByMap()
}
