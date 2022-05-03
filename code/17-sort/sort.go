// 使用sort.Slice函数进行优先级排序
package main

import (
	"fmt"
	"sort"
)

// 结构体，人
type People struct {
	name string
	age  int
}

func main() {
	arr := []People{}
	arr = append(arr, People{name: "张三", age: 20})
	arr = append(arr, People{name: "李四", age: 15})
	arr = append(arr, People{name: "王五", age: 30})
	fmt.Printf("原切片: %v\n", arr)
	// 将People切片按照每个人的年龄(age)从小到大进行排序
	sort.SliceStable(arr, func(i, j int) bool {
		// i, j 是索引
		// 排序的规则是：
		// 返回true，i优先；返回false，j优先
		return arr[i].age <= arr[j].age
	})
	fmt.Printf("排序之后的切片: %v\n", arr)
}
