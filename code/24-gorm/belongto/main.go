// belong to 一对一实体关联的使用
package main

import (
	"code/24-gorm/dao"
	"code/24-gorm/model"
	"fmt"
)

func main() {
	// 自动创建表
	dao.DB.AutoMigrate(&model.Class{}, &model.Student{})
	// 插入一个班级
	class1 := model.Class{}
	class1.ID = 2
	// dao.DB.Create(&class1)
	
	// 插入一个学生，并关联班级
	// student1 := model.Student{Name: "李四", Class: class1}
	// res := dao.DB.Create(&student1)
	// fmt.Printf("res.RowsAffected: %v\n", res.RowsAffected)
	
	// 查询学生
	students := []model.Student{}
	dao.DB.Find(&students)
	for _, student := range students {
		// 查找学生关联的班级信息
		dao.DB.Model(&student).Association("Class").Find(&student.Class)
		fmt.Printf("student: %v\n", student)
	}
}