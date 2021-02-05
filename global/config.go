package global

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	SensitiveWords  []string
	MessageQueueLen = 1024
	RedisAddr       string
)

func initConfig() {
	viper.SetConfigName("chatadmin")
	viper.AddConfigPath(RootDir + "/config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	SensitiveWords = viper.GetStringSlice("sensitive")
	MessageQueueLen = viper.GetInt("message-queue")
	RedisAddr = viper.GetString("redis")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		viper.ReadInConfig()

		SensitiveWords = viper.GetStringSlice("sensitive")
	})
}
