package log

import (
	"sync"
)

type Field struct {
	Key   string
	Value any
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

type Option struct {
	OutputPaths []string
}

// 单例模式+代理模式实现

// ILogger 日志管理器
type ILogger interface {
	newEntry(name string) IEntry
	SetDebug(debug bool)
}

var logger = struct {
	ILogger
	sync.Once
}{}

func Log() ILogger {
	logger.Do(func() {
		logger.ILogger = initLogger()
	})
	return logger
}

//todo 分为系统日志和用户日志

// NewEntry 创建一个新的日志记录器,name指的是日志记录器的对象
func NewEntry(name string) IEntry {
	return Log().newEntry(name)
}

// 加载的日志方法由配置文件决定
func initLogger() ILogger {
	//todo 直接多日志模块封装
	return zapLoggerInit()
}
