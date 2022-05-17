// gorm查询相关api
package main

import (
	"code/24-gorm/dao"
	"code/24-gorm/model"
	"fmt"
	"time"
)

// s16 Raw
func s16() {
	// 查询所有年龄小于7的宠物
	pets := []model.Pet{}
	dao.DB.Raw("select * from pets where age < ?", 7).Scan(&pets)
	for _, pet := range pets {
		fmt.Printf("pet: %v\n", pet)
	}
}

// s15 Distinct
func s15() {
	// 查询所有的宠物类型
	results := []string{}
	dao.DB.Model(&model.Pet{}).Distinct("type").Find(&results)
	for _, result := range results {
		fmt.Printf("result: %v\n", result)
	}
}

// s14 Group & Having
func s14() {
	// 统计结果结构体
	type Result struct {
		Date time.Time
		Total int
		Type string
	}
	results := []Result{}
	// 1. 使用结构体模型（查询大于5岁的每个类型的宠物只数）
	dao.DB.
	Model(&model.Pet{}).
	Select("count(*) as total, type").
	Where("age > ?", 5).
	Group("type").
	Find(&results)
	fmt.Printf("results: %v\n", results)
	// 2. 指定表名（查询每个类型的宠物只数，保留总数大于3的记录）
	dao.DB.
	Table("pets").
	Select("count(*) as total, type, created_at as date").
	Group("type").
	Having("total > ?", 3).
	Scan(&results)
	fmt.Printf("results: %v\n", results)
}

// s13 Limit & Offset
func s13() {
	var users []model.User
	// 从第一条记录开始（0是起始），取出1条数据
	dao.DB.Offset(1).Limit(1).Find(&users)
	for _, user := range users {
		fmt.Printf("user: %v\n", user)
	}
}

// s12 Order
func s12() {
	var users []model.User
	dao.DB.Order("id desc").Find(&users)
	for _, user := range users {
		fmt.Printf("user: %v\n", user)
	}
}

// s11 Select
func s11() {
	var user model.User
	dao.DB.Select("id", "username").First(&user)
	fmt.Printf("user: %v\n", user)
}

// s10 Or
func s10() {
	var users []model.User
	u := model.User{Username: "张三"}
	u.ID = 1
	dao.DB.Where("id = ?", 2).Or(&u).Find(&users)
	for _, user := range users {
		fmt.Printf("user: %v\n", user)
	}
}

// s9 Not
func s9() {
	var users []model.User
	dao.DB.Not("id IN ?", []int{1, 4}).Find(&users)
	for _, user := range users {
		fmt.Printf("user: %v\n", user)
	}
}

// s8 Struct Column
func s8() {
	var users []model.User
	// 结构体中两个参数是不匹配的，加以验证
	u := model.User{Username: "张三"}
	u.ID = 2
	dao.DB.Where(&u, "username").Find(&users)
	for _, user := range users {
		fmt.Printf("user: %v\n", user)
	}
}

// s7 Struct Map构造Where条件
func s7() {
	var users []model.User
	// 1. struct
	dao.DB.Where(&model.User{Username: "张三"}).Find(&users)
	for _, user := range users {
		fmt.Printf("user: %v\n", user)
	}
	// 2. map
	dao.DB.Where(map[string]interface{}{"username": "李四", "id": 2}).Find(&users)
	for _, user := range users {
		fmt.Printf("user: %v\n", user)
	}
}

// s6 使用字符串构造Where条件
func s6() {
	var users []model.User
	// 1. =
	fmt.Println("=")
	dao.DB.Where("id = ?", 1).Find(&users)
	for _, user := range users {
		fmt.Printf("user: %v\n", user)
	}
	// 2. <>
	fmt.Println("<>")
	dao.DB.Where("id <> ?", 1).Find(&users)
	for _, user := range users {
		fmt.Printf("user: %v\n", user)		
	}
	// 3. IN
	fmt.Println("IN")
	dao.DB.Where("id IN ?", []int{1, 4}).Find(&users)
	for _, user := range users {
		fmt.Printf("user: %v\n", user)
	}
	// 4. LIKE
	fmt.Println("LIKE")
	dao.DB.Where("username LIKE ?", "%三%").Find(&users)
	for _, user := range users {
		fmt.Printf("user: %v\n", user)
	}
	// 5. AND
	fmt.Println("AND")
	dao.DB.Where("id = ? AND username = ?", 1, "张三").Find(&users)
	for _, user := range users {
		fmt.Printf("user: %v\n", user)
	}
	// 6. TIME
	fmt.Println("TIME")
	dao.DB.Where("created_at < ?", time.Now()).Find(&users)
	for _, user := range users {
		fmt.Printf("user: %v\n", user)
	}
	// 7. BETWEEN
	fmt.Println("BETWEEN")
	dao.DB.Where("id BETWEEN ? AND ?", 2, 5).Find(&users)
	for _, user := range users {
		fmt.Printf("user: %v\n", user)
	}
}

// s5 Find
func s5() {
	var users []model.User
	res := dao.DB.Find(&users)
	fmt.Printf("res.RowsAffected: %v\n", res.RowsAffected)
	fmt.Printf("res.Error: %v\n", res.Error)
	// 打印所有的记录
	for _, user := range users {
		fmt.Printf("user: %v\n", user)
	}
}

// s4 Find One
func s4() {
	var user model.User
	res := dao.DB.Limit(1).Find(&user)
	fmt.Printf("user.ID: %v\n", user.ID)
	fmt.Printf("res.RowsAffected: %v\n", res.RowsAffected)
	fmt.Printf("res.Error: %v\n", res.Error)
	fmt.Printf("user: %v\n", user)
}

// s3 Last
func s3() {
	var user model.User
	res := dao.DB.Last(&user)
	fmt.Printf("user.ID: %v\n", user.ID)
	fmt.Printf("res.RowsAffected: %v\n", res.RowsAffected)
	fmt.Printf("res.Error: %v\n", res.Error)
	fmt.Printf("user: %v\n", user)
}

// s2 Take
func s2() {
	var user model.User
	res := dao.DB.Take(&user)
	fmt.Printf("user.ID: %v\n", user.ID)
	fmt.Printf("res.RowsAffected: %v\n", res.RowsAffected)
	fmt.Printf("res.Error: %v\n", res.Error)
	fmt.Printf("user: %v\n", user)
}

// s1 First
func s1() {
	user := model.User{}
	user.ID = 2
	res := dao.DB.First(&user)
	fmt.Printf("user.ID: %v\n", user.ID)
	fmt.Printf("res.RowsAffected: %v\n", res.RowsAffected)
	fmt.Printf("res.Error: %v\n", res.Error)
	fmt.Printf("user: %v\n", user)
}

func main() {
	fmt.Println("First:")
	s1()
	fmt.Println("==========================")
	fmt.Println("Take:")
	s2()
	fmt.Println("==========================")
	fmt.Println("Last:")
	s3()
	fmt.Println("==========================")
	fmt.Println("Find One:")
	s4()
	fmt.Println("==========================")
	fmt.Println("Find All:")
	s5()
	fmt.Println("==========================")
	fmt.Println("Where String:")
	s6()
	fmt.Println("==========================")
	fmt.Println("Where Struct & Map:")
	s7()
	fmt.Println("==========================")
	fmt.Println("Where Struct Column:")
	s8()
	fmt.Println("==========================")
	fmt.Println("Not:")
	s9()
	fmt.Println("==========================")
	fmt.Println("Or:")
	s10()
	fmt.Println("==========================")
	fmt.Println("Select:")
	s11()
	fmt.Println("==========================")
	fmt.Println("Order:")
	s12()
	fmt.Println("==========================")
	fmt.Println("Limit & Offset:")
	s13()
	fmt.Println("==========================")
	fmt.Println("Group & Having:")
	s14()
	fmt.Println("==========================")
	fmt.Println("Distinct:")
	s15()
	fmt.Println("==========================")
	fmt.Println("Raw:")
	s16()
	fmt.Println("==========================")
}
