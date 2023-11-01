package log

// IEntry 日志记录器
type IEntry interface {
	Log(level Level, msg string, args ...Field)
	Info(message string, args ...Field)
	Debug(msg string, args ...Field)
	Warn(msg string, args ...Field)
	Error(msg string, args ...Field)
	Fatal(msg string, args ...Field)
	Panic(msg string, args ...Field)
	Put()
}
