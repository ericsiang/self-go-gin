package zapLogger_test

import (
	"api/util/zapLogger"
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {
	config := &zapLogger.RotatelogsConfig{
		InfoLogPath:   "./info.log",
		ErrorLogPath:  "./error.log",
		MaxSize:       1024,
		RotationCount: 7,
		MaxAge:        24 * time.Hour,
		RotationTime:  24 * time.Hour,
	}

	logger, err := zapLogger.NewLogger(config)
	if err != nil {
		t.Fatalf("Failed to create new logger TestNewLogger() error : %v", err)
	}

	if logger == nil {
		t.Fatal("Expected logger to be not nil TestNewLogger()")
	}
}
