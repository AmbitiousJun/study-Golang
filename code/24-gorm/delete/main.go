package main

import (
	"code/24-gorm/dao"
	"code/24-gorm/model"
	"fmt"
)

// d1 Delete
func d1() {
	// 删除年龄为20的宠物
	dao.DB.Where("age = ?", 20).Delete(&model.Pet{})
}

// d2 DeleteBatch
func d2() {
	// 删除主键为9的记录
	dao.DB.Delete(&model.Pet{}, 9)
	// 删除主键为1，2，3的记录
	dao.DB.Delete(&model.Pet{}, []int{1, 2, 3})
}

// d3 Unscoped
func d3() {
	// 查询所有已经被删除的宠物
	pets := []model.Pet{}
	dao.DB.Unscoped().Where("deleted_at is not null").Find(&pets)
	for _, pet := range pets {
		fmt.Printf("pet: %v\n", pet)
	}
}

func main() {
	fmt.Println("Delete:")
	d1()
	fmt.Println("====================")
	fmt.Println("DeleteBatch:")
	d2()
	fmt.Println("====================")
	fmt.Println("Unscoped:")
	d3()
	fmt.Println("====================")
}
