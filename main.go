package main

import (
	"fmt"

	"github.com/chengjoey/chatadmin/router"

	"github.com/gin-gonic/gin"
)

var (
	addr   = ":9090"
	banner = "ChatRoom，start on：%s"
)

func init() {
}

func main() {
	fmt.Printf(banner, addr)
	server := gin.Default()
	router.RegisterHandle(server)
	server.Run("0.0.0.0:9090")
}
