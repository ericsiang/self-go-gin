package zapLogger

import (
	"errors"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type RotatelogsConfig struct {
	InfoLogPath   string
	ErrorLogPath  string
	MaxSize       int64
	RotationCount int64
	MaxAge        time.Duration
	RotationTime  time.Duration
}

type Logger struct {
	zapLogger *zap.Logger
}

func NewLogger(rotatelogsConfig *RotatelogsConfig) (*Logger, error) {
	zapLogger, err := initZap(rotatelogsConfig)
	if err != nil {
		return nil, err
	}

	zap.ReplaceGlobals(zapLogger) //使用全局logger(設定了在其他地方調用 zap.S() or zap.L() 才會生效)
	defer zapLogger.Sync()        // zap底层有缓冲。在任何情况下执行 defer logger.Sync() 是一个很好的习惯
	return &Logger{zapLogger: zapLogger}, nil
}

func (l *Logger) GetZapLogger() *zap.Logger {
	return l.zapLogger
}

func getFileRotatelogs(filePath string, rotatelogsConfig *RotatelogsConfig) (*rotatelogs.RotateLogs, error) {
	/*
		設定日誌輸出路徑，使用file-rotatelogs 進行切割
		MaxAge and RotationCount cannot be both set  兩者不能同時設置
		官方 github 上有說明建議使用 WithRotationCount，要將 MaxAge 設為 -1 比較保險
	*/

	logf, err := rotatelogs.New(filePath+".%Y%m%d",
		// rotatelogs.WithLinkName(filePath),      // 生成軟鏈，指向最新日誌文件
		rotatelogs.WithMaxAge(rotatelogsConfig.MaxAge),                     //保留舊日誌文件的最大天數
		rotatelogs.WithRotationTime(rotatelogsConfig.RotationTime),         //切割頻率為時間單位
		rotatelogs.WithRotationCount(uint(rotatelogsConfig.RotationCount)), //保留舊日誌文件最大保存個數
		rotatelogs.WithRotationSize(rotatelogsConfig.MaxSize),              //切割頻率為文件大小
		rotatelogs.WithHandler(rotatelogs.HandlerFunc(func(e rotatelogs.Event) {
			// 在這裡添加你的自定義操作
			if e.Type() == rotatelogs.FileRotatedEventType {
				// 這裡的代碼將在每次日誌切割時執行
				// e.(*rotatelogs.FileRotatedEvent).PrevFile() 是上一個日誌文件的路徑
				// e.(*rotatelogs.FileRotatedEvent).CurrentFile() 是當前日誌文件的路徑
			}

		})),
	)

	if err != nil {
		// log.Fatalln("zapLogger getFileRotatelogs()  err:", err)
		return nil, err
	}

	return logf, nil
}

func getNewCore(encoder zapcore.Encoder, logf *rotatelogs.RotateLogs, level zapcore.LevelEnabler) zapcore.Core {
	writeSyncer := zapcore.AddSync(logf)
	/*
		設定可以同時輸出到控制台和文件
		MultiWriteSyncer() => 一次輸出多個 WriteSyncer
		writeSyncer  => 輸出到文件
		zapcore.AddSync(os.Stdout) => 輸出到控制台
	*/
	return zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout)), level)
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func getEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",   //日誌時間的key
		LevelKey:       "level",  //日誌級別的key
		NameKey:        "logger", //日誌名的key
		MessageKey:     "msg",    //日誌消息的key
		CallerKey:      "caller", //日誌調用函數的key
		StacktraceKey:  "trace",
		LineEnding:     zapcore.DefaultLineEnding,      //日誌結尾分隔符 - 默認/n
		EncodeLevel:    zapcore.CapitalLevelEncoder,    //日志级别，默認小寫，這裡設定為大寫
		EncodeTime:     customTimeEncoder,              //自訂日誌輸出時間格式 - customTimeEncoder()
		EncodeDuration: zapcore.SecondsDurationEncoder, //執行消耗的時間轉換成浮點型的秒
		EncodeCaller:   zapcore.ShortCallerEncoder,     //路径编码器
	}

	return encoderConfig
}

func getEncoder() zapcore.Encoder {
	encoderConfig := getEncoderConfig()

	var encoder zapcore.Encoder
	switch gin.Mode() {
	case gin.DebugMode:
		encoder = zapcore.NewConsoleEncoder(encoderConfig) //使用 console 格式
	case gin.ReleaseMode:
		encoder = zapcore.NewJSONEncoder(encoderConfig) //使用 json 格式
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig) //使用 console 格式
	}

	return encoder
}

func getLevelEnabler() (infoLevel, errorLevel zap.LevelEnablerFunc) {
	infoLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})

	errorLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	return infoLevel, errorLevel
}
func initZap(rotatelogsConfig *RotatelogsConfig) (*zap.Logger, error) {
	encoder := getEncoder()

	//依不同級別寫入不同文件
	infoLevel, errorLevel := getLevelEnabler()

	// info level log file
	infoLogf, err := getFileRotatelogs(rotatelogsConfig.InfoLogPath, rotatelogsConfig)
	if err != nil {
		infoLogfError := errors.New("[zapLogger] initZap() infoLogf err : " + err.Error())
		// log.Fatalln("zapLogger initZap() infoLogf err:", err)
		return nil, infoLogfError
	}
	infoCore := getNewCore(encoder, infoLogf, infoLevel)

	// error level log file
	errorLogf, err := getFileRotatelogs(rotatelogsConfig.ErrorLogPath, rotatelogsConfig)
	if err != nil {
		errorLogfError := errors.New("[zapLogger] initZap() errorLogf err : " + err.Error())
		return nil, errorLogfError
	}
	errorCore := getNewCore(encoder, errorLogf, errorLevel)

	teeCore := zapcore.NewTee([]zapcore.Core{infoCore, errorCore}...)

	return zap.New(teeCore, zap.AddCaller(), zap.AddCallerSkip(1)), nil // zap.AddCaller()为顯示文件名和行號，可省略
}
