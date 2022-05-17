package main

import (
	"code/24-gorm/dao"
	"code/24-gorm/model"
	"fmt"
)

// u1 Save
func u1() {
	// 查询年龄为3的宠物，并将其名字改为狗狗
	pet := model.Pet{}
	res := dao.DB.Where("age = ?", 3).First(&pet)
	if res.Error != nil {
		fmt.Printf("res.Error: %v\n", res.Error)
	}
	fmt.Printf("pet: %v\n", pet)
	pet.Name = "狗狗"
	dao.DB.Save(&pet)
}

// u2 Update
func u2() {
	// 将年龄为4的宠物的名字修改为狗蛋
	dao.DB.Model(&model.Pet{}).Where("age = ?", 4).Update("name", "狗蛋")
}

// u3 Updates
func u3() {
	// 将id为7的宠物的年龄改为20,名字改为蛋蛋
	pet := model.Pet{}
	pet.ID = 7
	dao.DB.Model(&pet).Updates(&model.Pet{Age: 20, Name: "蛋蛋"})
}

func main() {
	fmt.Println("Save:")
	u1()
	fmt.Println("=========================")
	fmt.Println("Update:")
	u2()
	fmt.Println("=========================")
	fmt.Println("Updates:")
	u3()
	fmt.Println("=========================")
}
