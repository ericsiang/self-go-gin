package initialize

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var envPath = "env.yaml"

func initEnv(serverConfig *ServerConfig) {
	//viper 讀取環境變量的套件
	v := viper.New()

	//讀取配置文件
	v.SetConfigFile(envPath)

	//讀取配置信息
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	v.WatchConfig() // 監聽 env file
	v.OnConfigChange(func(e fsnotify.Event) {
		InitSetting()
		fmt.Println("Config file changed:", e.Name)
	})
}
