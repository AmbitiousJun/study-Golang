// 闭包的使用
package main

import (
	"fmt"
	"strings"
)

// 示例1：后缀函数生成器
// 给定一个文件后缀suffix，返回一个后缀生成函数
// 只需要传入一个字符串，就可以在字符串后面在拼接上给定的后缀
func MakeSuffixFunc(suffix string) func(string) string {
	return func(base string) string {
		if !strings.HasSuffix(base, suffix) {
			base += suffix
		}
		return base
	}
}

// 示例2：计数器
// 传入一个base值，每次参数计数器的运算
// 返回两个函数:
// add: 返回base + 形参
// sub: 返回base - 形参
func Cal(base int) (add func(int) int, sub func(int) int) {
	add = func(a int) int {
		return base + a
	}
	sub = func(a int) int {
		return base - a
	}
	return
}

func main() {
	// 示例1
	jpgSuffix := MakeSuffixFunc(".jpg")
	fmt.Println(jpgSuffix("test"))
	txtSuffix := MakeSuffixFunc(".txt")
	fmt.Println(txtSuffix("abc"))

	fmt.Println("====================")

	// 示例2
	af, sf := Cal(100)
	fmt.Println(af(30), sf(50))
	fmt.Println(af(40), sf(60))
	af, sf = Cal(100)
	fmt.Println(af(3), sf(4))
	fmt.Println(af(5), sf(6))
}
