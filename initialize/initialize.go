package initialize

import (
	"go.uber.org/zap"
	unitrans "github.com/go-playground/universal-translator"
)

var (
	ServerEnv = ServerConfig{}
	Logger *zap.Logger

	Trans unitrans.Translator
)

func InitConfig() {
	initEnv(&ServerEnv)
	initLogger(Logger) 
	zap.S().Info("配置信息 : " , ServerEnv)
	initMysql()
	// vaildate 中文化
	if err := 	initValidateLang("zh") ; err != nil { 
		zap.S().Error("init trans failed, err:%v\n", err.Error())
		return
	}
}
