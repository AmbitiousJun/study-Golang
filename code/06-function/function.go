// 函数的使用
package main

import (
	"fmt"
)

func main() {
	max := max(2, 3)
	fmt.Println(max)
	fmt.Println("======================")
	a, b := swap("Google", "Tencent")
	fmt.Println(a, b)
	fmt.Println("======================")
	c, d := 200, 300
	fmt.Printf("交换前c=%d, d=%d\n", c, d)
	swapInt(&c, &d)
	fmt.Printf("交换后c=%d, d=%d\n", c, d)
}

// 将传递过来的两个引用进行交换
func swapInt(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

// 将传递来的两个字符串交换并返回
func swap(a, b string) (string, string) {
	return b, a
}

// 求两数的最大值，接收两个整型数，返回两个数中较大数
func max(a, b int) int {
	// 定义局部变量
	var res int
	if a < b {
		res = b
	} else {
		res = a
	}
	return res
}
