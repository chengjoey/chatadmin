package main

import (
	"embed"
	"fmt"

	"github.com/chengjoey/chatadmin/global"
	"github.com/chengjoey/chatadmin/router"
	"github.com/gin-gonic/gin"
)

var (
	banner = "ChatRoom，start on "
	//go:embed template
	embededFiles embed.FS
	//go:embed config/chatadmin.yaml
	configData []byte
)

func init() {
	global.Init(configData)
}

func main() {
	fmt.Println(banner, global.ServerAddr)
	server := gin.Default()
	router.RegisterHandle(server, embededFiles)
	server.Run(global.ServerAddr)
}
