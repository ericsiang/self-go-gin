package initialize

import (
	"go.uber.org/zap"
)

var (
	serverEnv = &ServerConfig{}
)

func InitSetting() {
	initLogger()
	initEnv(serverEnv)
	zap.S().Info("配置信息 : ", serverEnv)
	initMysql()
	initRedis()
	// vaildate 中文化
	if err := initValidateLang("zh"); err != nil {
		zap.S().Error("init trans failed, err:%v\n", err.Error())
		panic(err)
	}
}

func GetServerEnv() *ServerConfig {
	return serverEnv
}
