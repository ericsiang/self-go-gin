package initialize

import (
	"github.com/spf13/viper"
)

func initEnv() {
	//viper 讀取環境變量的套件
	v := viper.New()

	//讀取配置文件
	v.SetConfigFile("config.yaml")
	//讀取配置信息
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
}
