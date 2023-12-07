package initialize

import (
	 "api/common/common_const"
	"api/util/zapLogger"
	"log"

	"go.uber.org/zap"
)

func GetZapLogger() *zap.Logger {
	rotatelogsConfig := &zapLogger.RotatelogsConfig{
		InfoLogPath: "log/info/info.log",
		ErrorLogPath: "log/error/error.log",
		MaxSize: common_const.ZapLoggerMaxSize,
		RotationCount: common_const.ZapLoggerMaxCounts,
		MaxAge: common_const.ZapLoggerMaxAge,
		RotationTime: common_const.ZapLoggerRotationTime,
	}
	
	zapLogger, err := zapLogger.NewLogger(rotatelogsConfig)
	if err != nil {
		log.Fatalln("[logger] GetZapLogger() err : ", err)
	}

	return zapLogger.GetZapLogger()
}
