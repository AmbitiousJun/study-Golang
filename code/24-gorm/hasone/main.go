// has one实现user和creditCard的一对一关联
package main

import (
	"code/24-gorm/dao"
	"code/24-gorm/model"
	"fmt"
)

func main() {
	// 自动创建表
	dao.DB.AutoMigrate(&model.User{}, &model.CreditCard{})
	// 插入一个用户，同时存入身份证号
	// cc := model.CreditCard{Number: "440583223909043902923"}
	// u := model.User{Username: "憨憨", Password: "123456", CreditCard: cc}
	// res := dao.DB.Create(&u)
	// fmt.Printf("res.RowsAffected: %v\n", res.RowsAffected)

	// 查询用户，并封装信息
	u1 := model.User{}
	dao.DB.Where("username = ?", "憨憨").First(&u1)
	dao.DB.Model(&u1).Association("CreditCard").Find(&u1.CreditCard)
	fmt.Printf("u1: %v\n", u1)
}
