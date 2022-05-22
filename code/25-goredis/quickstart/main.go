// go-redis入门程序
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// 初始化redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		// 使用默认的数据库
		DB: 0,
	})
	// 设置一个string类型的值
	err := rdb.Set(ctx, "hello", "world", 0).Err()
	if err != nil {
		log.Fatal(err)
	}
	// 取出一个存在的string类型的值
	res, err := rdb.Get(ctx, "hello").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("key: hello, val:", res)
	// 取出一个不存在的string类型的值
	res, err = rdb.Get(ctx, "hello1").Result()
	if err == redis.Nil {
		log.Fatal("不存在key:", "hello1")
	} else if err != nil {
		log.Fatal(err)
	}
	fmt.Println("key: hello1, val:", res)
}
