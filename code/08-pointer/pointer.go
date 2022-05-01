// 指针的使用
package main

import "fmt"

func main() {
	// 1 定义一个整型变量和一个整型指针，用指针指向该变量的内存地址
	var a int
	var ap *int
	a = 1
	ap = &a
	fmt.Printf("变量a的地址为: %x\n", &a)
	fmt.Printf("指针ap的值为: %x\n", ap)
	fmt.Printf("指针ap指向的地址的内容为: %d\n", *ap)
}
