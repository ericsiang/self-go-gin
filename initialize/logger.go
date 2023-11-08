package initialize

import (
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func getWriteSyncer(filePath string) zapcore.WriteSyncer {
	//使用file-rotatelogs 進行切割，每天一个文件
	logf, _ := rotatelogs.New(filePath+".%Y%m%d",
		rotatelogs.WithLinkName(filePath),        // 生成軟鏈，指向最新日誌文件
		rotatelogs.WithMaxAge(7*24*time.Hour),    // 保存7天
		rotatelogs.WithRotationTime(time.Minute), //切割頻率
	)

	return zapcore.AddSync(logf)
}

func getNewCore(encoder zapcore.Encoder, filePath string, level zapcore.LevelEnabler) zapcore.Core {
	writeSyncer := getWriteSyncer(filePath)
	return zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout)), level)
}

func initZap() *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",   //日誌時間的key
		LevelKey:       "level",  //日誌級別的key
		NameKey:        "logger", //日誌名的key
		MessageKey:     "msg",    //日誌消息的key
		CallerKey:      "caller", //日誌調用函數的key
		StacktraceKey:  "trace",
		LineEnding:     zapcore.DefaultLineEnding,      //日誌結尾分隔符 - 默認/n
		EncodeLevel:    zapcore.CapitalLevelEncoder,  //日志级别，默認小寫
		EncodeTime:     CustomTimeEncoder,              //自訂日誌輸出時間格式 - CustomTimeEncoder()
		EncodeDuration: zapcore.SecondsDurationEncoder, //執行消耗的時間轉換成浮點型的秒
		EncodeCaller:   zapcore.ShortCallerEncoder,     //路径编码器
	}
	// encoder := zapcore.NewJSONEncoder(encoderConfig) //使用 json 格式
	encoder := zapcore.NewConsoleEncoder(encoderConfig) //使用 json 格式

	//依不同級別寫入不同文件，並設定日誌切割規則
	cores := [...]zapcore.Core{
		//info level log
		getNewCore(encoder, "./log/info/info.log", zapcore.InfoLevel),
		//error level log
		getNewCore(encoder, "./log/error/error.log", zapcore.ErrorLevel),
		//debug level log
		getNewCore(encoder, "./log/debug/debug.log", zapcore.DebugLevel),
		//warn level log
		getNewCore(encoder, "./log/warn/warn.log", zapcore.WarnLevel),
	}

	teeCore := zapcore.NewTee(cores[:]...)

	return zap.New(teeCore, zap.AddCaller(), zap.AddCallerSkip(1))
}

var logger *zap.Logger

func initLogger() {
	//初始化zap日志
	logger = initZap()
	defer logger.Sync() // zap底层有缓冲。在任何情况下执行 defer logger.Sync() 是一个很好的习惯

	zap.ReplaceGlobals(logger) //使用全局logger(設定了在其他地方調用 zap.S() or zap.L() 才會生效)
}

func GetLogger() *zap.Logger {
	return logger
}
