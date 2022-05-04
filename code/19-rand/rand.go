// 使用rand包产生随机数
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 200
	arr := []int{}
	// 添加随机数种子
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 10; i++ {
		arr = append(arr, rand.Intn(n))
	}
	fmt.Println(arr)
}
