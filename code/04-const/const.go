// 常量的使用
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 1 常量完整的定义形式
	const WIDTH int = 10
	const HEIGHT int = 20
	area := WIDTH * HEIGHT
	fmt.Println(area)
	// 2 使用并行赋值常量
	const a, b, c = 1, false, "abc"
	fmt.Println(a, b, c)
	// 3 常量用作枚举
	const (
		// 未知
		Unknow = 0
		// 男性
		Female = 1
		// 女性
		Male = 2
	)
	fmt.Println(Unknow, Female, Male)
	// 4 在常量中使用内置函数计算表达式的值
	const (
		d = "abc"
		e = len(d)
		f = unsafe.Sizeof(a)
	)
	fmt.Println(d, e, f)
	fmt.Println("=====================")
	iota_test()
}

func iota_test() {
	// iota常量用作枚举
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)
	// 常量枚举的简写
	const (
		d = iota
		e
		f
	)
	fmt.Println(d, e, f)
}
