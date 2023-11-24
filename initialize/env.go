package initialize

import (
	"github.com/spf13/viper"
)



func initEnv(serverConfig *ServerConfig) {
	//viper 讀取環境變量的套件
	v := viper.New()

	//讀取配置文件
	v.SetConfigFile("env.yaml")
	//讀取配置信息
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
}
