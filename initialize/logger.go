package initialize

import (
	"api/common/common_const"
	"api/util/zap_logger"
	"log"

	"go.uber.org/zap"
)

func GetZapLogger() *zap.Logger {
	rotatelogsConfig := &zap_logger.RotatelogsConfig{
		InfoLogPath:   "log/info/info_%Y-%m-%d.log",
		ErrorLogPath:  "log/error/error_%Y-%m-%d.log",
		MaxSize:       common_const.ZapLoggerMaxSize,
		RotationCount: common_const.ZapLoggerMaxCounts,
		MaxAge:        common_const.ZapLoggerMaxAge,
		RotationTime:  common_const.ZapLoggerRotationTime,
	}

	zapLogger, err := zap_logger.NewLogger(rotatelogsConfig)
	if err != nil {
		log.Fatalln("[logger] GetZapLogger() err : ", err)
	}

	return zapLogger.GetZapLogger()
}
