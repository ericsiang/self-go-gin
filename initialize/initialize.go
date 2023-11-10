package initialize

import (

	"go.uber.org/zap"
)

var (
	ServerEnv = ServerConfig{}
	Logger *zap.Logger
)

func InitConfig() {
	initEnv(&ServerEnv)
	initLogger(Logger)

	zap.S().Info("配置信息 : " , ServerEnv)
	initMysql()
}
