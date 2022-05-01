// 递归的使用
package main

import "fmt"

func main() {
	ans := factorial(15)
	fmt.Printf("15的阶乘为: %d\n", ans)
}

// 递归求解一个数的阶乘
func factorial(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
