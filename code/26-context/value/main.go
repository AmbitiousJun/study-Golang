// WithValue的使用
package main

import (
	"context"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

const KEY = "trace_id"

func routine2(ctx context.Context) {
	printLog(ctx.Value(KEY).(uuid.UUID), "routine2输出了一条日志")
}

func routine1(ctx context.Context) {
	printLog(ctx.Value(KEY).(uuid.UUID), "routine1输出了一条日志")
	go routine2(ctx)
}

// printLog 输出日志到控制台上
func printLog(traceId uuid.UUID, text string) {
	fmt.Printf("%s|info|trace_id=%v|%s\n", time.Now().Format("2006-01-02 15:04:05"), traceId, text)
}

func main() {
	// 根context
	ctx := context.Background()
	// 衍生withValue
	valueCtx := context.WithValue(ctx, KEY, uuid.NewV4())
	go routine1(valueCtx)
	time.Sleep(time.Second * 5)
}	
