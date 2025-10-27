package zap_log

import (
	"log"
	"self_go_gin/common/common_const"
	"self_go_gin/util/zap_logger"

	"go.uber.org/zap"
)

func GetZapLogger(logPath string) *zap.Logger {
	rotatelogsConfig := &zap_logger.RotatelogsConfig{
		InfoLogPath:   logPath + "info/info_%Y-%m-%d.log",
		ErrorLogPath:  logPath + "error/error_%Y-%m-%d.log",
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
