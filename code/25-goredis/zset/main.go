// zset类型的使用
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
	items := []*redis.Z{
		{Score: 50, Member: "C++"},
		{Score: 75, Member: "Java"},
		{Score: 25, Member: "PHP"},
		{Score: 100, Member: "Golang"},
		{Score: 0, Member: "Python"},
	}
	ctx := context.Background()
	// 设置值
	err = rdb.ZAdd(ctx, "test:zset", items...).Err()
	if err != nil {
		log.Fatal(err)
	}

	// 取出集合中按分数从小到大排的前三位
	res, err := rdb.ZRange(ctx, "test:zset", 0, 2).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res: %v\n", res)

	// 取出集合中按分数从大到小排的前三位
	res, err = rdb.ZRevRange(ctx, "test:zset", 0, 2).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res: %v\n", res)

	// 取出集合中的所有元素（按分数从小到大排）
	res, err = rdb.ZRange(ctx, "test:zset", 0, -1).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res: %v\n", res)
}
