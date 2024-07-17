package initialize

import (
	"api/util/jwt_secret"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	serverEnv = &ServerConfig{}
)

func InitSetting() {
	initEnv(serverEnv)
	zap.S().Info("配置信息 : ", serverEnv)
	gin.SetMode(serverEnv.APP_Mode)
	initMysql()
	initRedis()
	jwt_secret.SetJwtSecret(GetServerEnv().JwtSecret)
	// vaildate 中文化
	if err := initValidateLang("zh"); err != nil {
		zap.S().Error("init trans failed, err:%v\n", err.Error())
		panic(err)
	}
}

func GetServerEnv() *ServerConfig {
	return serverEnv
}
