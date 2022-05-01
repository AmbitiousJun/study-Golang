// 并发使用
package main

import (
	"fmt"
	// "time"
)

// 每隔1s打印字符串一次
// func say(str string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(time.Second)
// 		fmt.Println(str)
// 	}
// }

// 切片求和
// func sum(slice []int, c chan int) {
// 	sum := 0
// 	for _, v := range slice {
// 		sum += v
// 	}
// 	// 将求和结果写入通道中
// 	fmt.Println(sum)
// 	c <- sum
// }

// 生成n个斐波那契数，存入通道中
func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	// 1 go并发编程的基本使用
	// go say("hello")
	// say("world")

	// 2 使用不带缓冲区的通道对一个数组进行求和
	// arr := []int{3, 9, -3, 2, 7, -6}
	// // 创建一个不带缓存的通道
	// c := make(chan int)
	// go sum(arr[:3], c)
	// go sum(arr[3:], c)
	// x, y := <-c, <-c
	// fmt.Println(x, y, x+y)

	// 3 使用带缓冲区的通道
	// ch := make(chan int, 100)
	// // 往通道中一次性存放两个数据，不会阻塞
	// ch <- 1
	// ch <- 2
	// // 将通道中的数据一次性取出来
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// 4 生成斐波那契数到通道中，并遍历打印
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)
	for v := range ch {
		fmt.Println(v)
	}
}
