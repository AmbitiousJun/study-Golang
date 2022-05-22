// 操作string类型数据
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

// 初始化连接
func initDB() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return client, err
}

func main() {
	rdb, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	// 设置值，过期时间为10分钟
	err = rdb.Set(ctx, "test:name", "张三", time.Minute*10).Err()
	if err != nil {
		log.Fatal(err)
	}
	// 获取值
	res, err := rdb.Get(ctx, "test:name").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res: %v\n", res)
}
