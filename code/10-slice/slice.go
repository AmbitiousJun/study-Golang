// 切片初始化
package main

import "fmt"

func main() {
	// 1 与数组相同的初始化方式
	slice1 := []int{1, 2, 3}
	printSlice(slice1)
	// 2 初始化切片为一个数组的引用
	arr := [5]int{1, 2, 3, 4, 5}
	slice2 := arr[:]
	printSlice(slice2)
	// 3 以数组中某个范围内的元素作为初始值初始化切片
	slice3 := arr[1:3]
	printSlice(slice3)

	fmt.Println("=============================")

	// 增大切片容量
	var numbers1 []int
	// 向切片中追加1个元素
	numbers1 = append(numbers1, 1)
	// 向切片中追加多个元素
	numbers1 = append(numbers1, 2, 3, 4)
	printSlice(numbers1)
	// 创建一个容量为原来的两倍的切片
	numbers2 := make([]int, len(numbers1), 2*len(numbers1))
	// 将原切片的数据拷贝到新切片上
	copy(numbers2, numbers1)
	printSlice(numbers2)
}

// 打印切片长度、容量、内容
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
