package env

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var envFile = "env.yaml"

// InitEnv 初始化環境變量
/*
envPath: 環境變量文件所在路徑
serverConfig: 服務器配置實例
reload_func: 配置文件改變時的回調函數
*/
func InitEnv(envPath string, serverConfig *ServerConfig, reloadFunc func()) {
	//viper 讀取環境變量的套件
	v := viper.New()
	fmt.Printf("讀取配置文件: %s\n", envPath+envFile)
	//讀取配置文件
	v.SetConfigFile(envPath + envFile)

	//讀取配置信息
	if err := v.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "配置文件讀取錯誤:", err)
		panic(err)
	}
	if err := v.Unmarshal(&serverConfig); err != nil {
		fmt.Fprintln(os.Stderr, "配置文件解析錯誤:", err)
		panic(err)
	}
	v.WatchConfig() // 監聽 env file
	v.OnConfigChange(func(e fsnotify.Event) {
		reloadFunc()
		fmt.Println("Config file changed:", e.Name)
	})
}
