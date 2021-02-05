package main

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	// defer rdb.Close()
}

func main() {
	// fmt.Println(rdb.HGet("users", "z").Result())
	fmt.Println(rdb.XRange("user:zc:channel", "-", "+").Result())
	// status := rdb.XGroupCreate("stm", "stm-group", "$")
	// fmt.Println(status.Result())
	// args := &redis.XReadArgs{Streams: []string{"stm", "$"}, Block: time.Duration(0), Count: 2}
	// for {
	// 	stm, err := rdb.XRead(args).Result()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	s := stm[0]
	// 	for _, msg := range s.Messages {
	// 		fmt.Println(msg.Values)
	// 	}
	// }
}
