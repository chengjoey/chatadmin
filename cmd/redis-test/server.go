package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-redis/redis/v7"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	// defer rdb.Close()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := map[string]interface{}{"content": scanner.Text()}
		arg := &redis.XAddArgs{Stream: "stm", ID: "*", Values: data}
		fmt.Println(rdb.XAdd(arg).Result())
	}
}
