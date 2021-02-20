package global

import (
	"bytes"

	"github.com/spf13/viper"
)

var (
	RedisAddr  string
	ServerAddr string
)

func initConfig(byteData []byte) {
	viper.SetConfigType("yaml")
	r := bytes.NewReader(byteData)
	// viper.AddConfigPath(RootDir + "/config")

	if err := viper.ReadConfig(r); err != nil {
		panic(err)
	}
	RedisAddr = viper.GetString("redis")
	ServerAddr = viper.GetString("server")
}
