package dao

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// init 初始化数据库连接
func init() {
	dsn := "root:cyj070723@tcp(localhost:3306)/go_db?charset=utf8mb4&parseTime=true"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = conn
}