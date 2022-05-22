// hash类型的使用
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
	// 设置Hash方式1
	err = rdb.HSet(ctx, "test:hash", "key1", "value1", "key2", "value2").Err()
	if err != nil {
		log.Fatal(err)
	}
	// 设置Hash方式2
	err = rdb.HSet(ctx, "test:hash", map[string]string{"key3": "value3", "key4": "value4"}).Err()
	if err != nil {
		log.Fatal(err)
	}
	// 获取Hash中某个key的值
	res, err := rdb.HGet(ctx, "test:hash", "key1").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res: %v\n", res)
	// 获取Hash中的所有值
	res1, err := rdb.HGetAll(ctx, "test:hash").Result()
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range res1 {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
}
