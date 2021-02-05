package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"nhooyr.io/websocket/wsjson"

	"nhooyr.io/websocket"
)

func main() {
	ctx := context.Background()

	token, to := os.Args[1], os.Args[2]
	url := fmt.Sprintf("ws://localhost:9090/v1/chat?client=%s", to)
	options := &websocket.DialOptions{HTTPHeader: http.Header{"Authorization": []string{token}}}
	c, _, err := websocket.Dial(ctx, url, options)
	if err != nil {
		panic(err)
	}
	defer c.Close(websocket.StatusInternalError, "内部出错")
	data := map[string]string{
		"content": "test send msg",
	}

	err = wsjson.Write(ctx, c, data)
	if err != nil {
		panic(err)
	}
	// var v interface{}
	var v map[string]interface{}
	go func() {
		for {
			err = wsjson.Read(ctx, c, &v)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Accept from server: %v\n", v)
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var data map[string]string
		_ = json.Unmarshal([]byte(strings.Replace(scanner.Text(), "\n", "", 0)), &data)
		wsjson.Write(ctx, c, data)
	}
	c.Close(websocket.StatusNormalClosure, "")
}
