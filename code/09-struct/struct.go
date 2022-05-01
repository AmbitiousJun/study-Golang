// 结构体的使用
package main

import "fmt"

// 图书结构
type Book struct {
	id    int
	title string
}

func main() {
	// 声明结构体1
	book1 := Book{1, "计算机网络"}
	// 声明结构体2
	book2 := Book{title: "程序员代码面试指南", id: 2}
	fmt.Println(book1, book2)
}
