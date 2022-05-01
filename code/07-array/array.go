// 数组的使用
package main

import "fmt"

func main() {
	// 1 常规初始化方法
	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	// 2 使用字面量快速初始化
	arr2 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr2)
	// 3 让编译器自动推断数组的长度
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr3)
	// 4 通过指定下标初始化元素（前提是数组已经指定了长度）
	arr4 := [5]int{1: 3, 3: 8}
	fmt.Println(arr4)

	fmt.Println("=================================")

	// 二维数组
	// 1 在声明的时候给整个数组直接赋值
	arr5 := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		// 最后一行的末尾必须加上逗号，否则编译器会报错
		{9, 10, 11, 12},
	}
	fmt.Println(arr5)
	// 2 声明数组之后，使用`append`函数进行赋值
	// 这里不能指定数组的大小，否则编译器会报错，具体原因日后研究
	arr6 := [][]int{}
	arr6 = append(arr6, []int{1, 2, 3})
	arr6 = append(arr6, []int{4, 5, 6})
	fmt.Println(arr6)

	fmt.Println("==================================")

	// 定义函数`getAverage`，接收一个数组，并计算数组的平均值
	arr7 := []int{1, 2, 5, 8, 1}
	ans := getAverage(arr7)
	fmt.Printf("数组的平均值为：%.3f", ans)
}

// 计算一个数组的平均值
func getAverage(arr []int) (ans float32) {
	if arr == nil {
		return
	}
	// 获取数组的大小
	n := len(arr)
	for i := 0; i < n; i++ {
		ans += float32(i)
	}
	ans /= float32(n)
	return
}
