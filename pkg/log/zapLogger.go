package log

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

//第一版设计思想
//1.日志输出到文件和命令窗口(全局均是如此)

// zapLogger is the zap logger
type zapLogger struct {
	//Logger    *zap.Logger
	once sync.Once
	zap.Config
}

var zapLoggerG = new(zapLogger)

func zapLoggerInit() *zapLogger {
	zapLoggerG.once.Do(
		func() {
			// 默认为生产环境
			zapLoggerG.Config = zap.NewProductionConfig()
		})
	return zapLoggerG
}

func zapLevel(level Level) zapcore.Level {
	//todo DPanicLevel 设置
	switch level {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	//case DPanicLevel:
	//	return zapcore.DPanicLevel
	case PanicLevel:
		return zapcore.PanicLevel
	case FatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func (z *zapLogger) setDebug(debug bool) {
	if debug {
		//开发环境下的日志

		// Debug 级别
		zapLoggerG.Config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		// 开启开发模式
		zapLoggerG.Config.Development = true
		// 返回一个用于开发环境的编码器配置。
		zapLoggerG.Config.EncoderConfig = zap.NewDevelopmentEncoderConfig()
		// 启用堆栈跟踪
		zapLoggerG.DisableCaller = false
		zapLoggerG.DisableStacktrace = false
		// 日志直接打印到控制台
		zapLoggerG.Config.OutputPaths = []string{"stdout",
			"./log/zap.log"}
	} else {
		var err error
		//生产环境下的日志
		//todo 生产环境下的日志配置
		//zapLoggerG.Logger, err = zap.NewProduction()
		zapLoggerG.Config = zap.NewProductionConfig()
		zapLoggerG.Config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zapLoggerG.DisableCaller = true
		zapLoggerG.DisableStacktrace = true
		if err != nil {
			panic(errors.Wrapf(err, "log init failed"))
		}

		//todo 日志输出路径
		zapLoggerG.Config.OutputPaths = []string{"stdout",
			"./log/zap.log"}
	}

}

// todo 加上对象池
func (z *zapLogger) newEntry(name string) IEntry {
	var err error
	entry := new(zapEntry)
	z.InitialFields = map[string]interface{}{
		"module": name,
	}
	entry.Logger, err = z.Config.Build()

	if err != nil {
		//todo 日志初始化失败
		panic(errors.Wrapf(err, "log entry init failed"))
	}
	////entry.Logger = entry.Logger.Named(name)
	//entry.Logger.With(
	//	zap.Namespace(name),
	//)
	return entry
}
