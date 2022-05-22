// set类型的使用
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
	// 往set中添加值
	err = rdb.SAdd(ctx, "test:set", 1, 2, 3, 4, 5, 6, 7).Err()
	if err != nil {
		log.Fatal(err)
	}
	// 查看set中是否存在某个值
	fmt.Printf("is number 3 in set: %v\n", rdb.SIsMember(ctx, "test:set", 3).Val())
	fmt.Printf("is number 8 in set: %v\n", rdb.SIsMember(ctx, "test:set", 8).Val())
}
