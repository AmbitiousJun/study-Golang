// WithCancel的使用
package main

import (
	"context"
	"fmt"
	"time"
)

func speak(ctx context.Context) {
	// 无限循环，每秒输出一行提示信息
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			fmt.Println("已闭嘴")
			return
		default:
			fmt.Println("正在讲话...")
		}
	}
}

func main() {
	ctx := context.Background()
	ctxCancel, cancel := context.WithCancel(ctx)
	// 启动goroutine
	go speak(ctxCancel)
	// 主线程等待10秒之后，主动终止goroutine的执行
	time.Sleep(time.Second * 10)
	fmt.Println("主线程让goroutine闭嘴")
	cancel()
	time.Sleep(time.Second)
}
