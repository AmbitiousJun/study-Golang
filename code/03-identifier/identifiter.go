package main

import "fmt"

func main() {
	identify()
	fmt.Println("================")
	definition()
	fmt.Println("================")
	definition1()
	fmt.Println("================")
	definition2()
	fmt.Println("================")
	definition3()
}

func definition3() {
	// 多变量声明
	var a, b, c int
	a, b, c = 1, 2, 3
	fmt.Println(a, b, c)

	var d, e, f = 4, 5, 6
	fmt.Println(d, e, f)

	g, h, i := 7, 8, 9
	fmt.Println(g, h, i)

	// 使用因式分解的写法一般用于声明全局变量
	var (
		j int
		k string
	)
	j = 10
	k = "hahaha"
	fmt.Println(j, k)
}

func definition2() {
	// 使用:=进行变量声明
	intVal := 1
	stringVal := "Ambitious"
	fmt.Println(intVal, stringVal)
}

func definition1() {
	// 根据值自动判断变量类型
	var a = true
	fmt.Println(a)
}

func definition() {
	// 变量声明
	// 1 手动赋值初始值
	var name = "ZhangSan"
	// 2 使用默认零值
	var a int
	var b bool
	var c float64
	var d string
	fmt.Println(name, a, b, c, d)
}

func identify() {
	// 声明变量示例
	var name string = "Ambitious"
	fmt.Println(name)
	var code1, code2 int = 1, 2
	fmt.Println(code1, code2)
}
