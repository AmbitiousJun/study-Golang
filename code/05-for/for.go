// for循环的使用
package main

import "fmt"

func main() {
	numbers := [6]int{1, 2, 3, 4, 5}
	n := 100
	arr := make([]int, n)
	fmt.Printf("数组arr的长度为：%d\n", len(arr))
	for index, number := range numbers {
		fmt.Printf("第 %d 位的数为 %d\n", index, number)
	}
}
