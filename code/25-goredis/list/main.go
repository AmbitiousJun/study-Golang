// list类型的使用
package main

import (
	"context"
	"fmt"
	"log"

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
	// 往list的右侧插入数据
	rdb.RPush(ctx, "test:list", 1, 2, 3, 4, 5)
	// 从list的左侧插入数据
	rdb.LPush(ctx, "test:list", 11, 22, 33, 44, 55)
	length := rdb.LLen(ctx, "test:list").Val()
	for i := int64(0); i < length / 2; i++ {
		fmt.Printf("rpop: %v\n", rdb.RPop(ctx, "test:list").Val())
		fmt.Printf("lpop: %v\n", rdb.LPop(ctx, "test:list").Val())
	}
}
