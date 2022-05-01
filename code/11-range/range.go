// range的使用
package main

import "fmt"

func main() {
	// 定义一个数组
	arr := [5]int{1, 2, 3, 4, 5}
	// 场景1：求和，只需要用到数本身，所以第一个参数可抛弃
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum)
	// 场景2：查找，找到数组中某个值位于哪一索引下
	for index, item := range arr {
		if item == 3 {
			fmt.Printf("数字3对应的索引为: %d\n", index)
		}
	}
	// 场景3：遍历集合
	kvs := map[string]string{"chen": "junjun", "zhang": "san"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// 场景4：按照Unicode遍历字符串
	for i, v := range "go" {
		fmt.Printf("i: %d, v: %c\n", i, v)
	}
}
