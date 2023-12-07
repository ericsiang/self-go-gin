package common_const

import "time"

const (
	ZapLoggerMaxSize      = 1 * 1024 * 1024 // log 大保存容量，單位為Byte
	ZapLoggerMaxCounts    = 2               // log 最大保存個數，單位為個數
	ZapLoggerMaxAge       = -1   // log 最大保存時間，單位為時間單位，不使用時請設定為 -1
	ZapLoggerRotationTime = 24 * time.Hour  // log 切割頻率，單位為時間單位
)
