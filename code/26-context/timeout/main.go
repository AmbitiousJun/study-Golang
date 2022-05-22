// WithTimeout的使用
package main

import (
	"context"
	"fmt"
	"time"
)

func deal(ctx context.Context) {
	for i := 0; i <= 10; i++ {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("程序终止")
			return
		default:
			fmt.Println("current i:", i)
		}
	}
}

func main() {
	ctx := context.Background()
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	deal(ctxTimeout)
}
