package log

import (
	"sync"
)

// 单例模式+代理模式实现

// ILogger 日志管理器
type ILogger interface {
	NewEntry(name string, options ...Options) IEntry
}

// IEntry 日志记录器
type IEntry interface {
	Log(level Level, msg string, args ...Field)
	Info(message string, args ...Field)
	Debug(msg string, args ...Field)
	Warn(msg string, args ...Field)
	Error(msg string, args ...Field)
	Fatal(msg string, args ...Field)
	Panic(msg string, args ...Field)
}

var logger ILogger
var onceLogger sync.Once

func Log() ILogger {
	onceLogger.Do(func() {
		logger = initLogger()
	})
	return logger
}

// 加载的日志方法由配置文件决定
func initLogger() ILogger {
	//todo 直接多日志模块封装
	return ZapLogger()
}

type Field struct {
	Key   string
	Value any
}
type Options struct {
}

// Level type
type Level uint32

const (
	ErrorLevel Level = iota
	WarnLevel
	InfoLevel
	DebugLevel
	PanicLevel
	FatalLevel
)
