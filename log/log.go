package log

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	instance             *zap.SugaredLogger
	output               zapcore.WriteSyncer
	defaultEncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "TIME",
		LevelKey:       "LEVEL",
		NameKey:        "LOGGER",
		CallerKey:      "LINE",
		MessageKey:     "MESSAGE",
		StacktraceKey:  "STACKTRACE",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 大写编码器，debug 模式时带颜色
		EncodeTime:     zapcore.ISO8601TimeEncoder,       // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
)

// Init log instance
func Init(appName, logPath string, debug bool) {
	ws := []zapcore.WriteSyncer{zapcore.AddSync(os.Stdout)}

	if logPath := logPath; logPath != "" {
		ws = append(ws, zapcore.AddSync(&lumberjack.Logger{
			Filename:   logPath,
			MaxSize:    500,
			MaxBackups: 3,
			MaxAge:     28,
			Compress:   true,
		}))
	}

	output = zapcore.NewMultiWriteSyncer(ws...)
	encoder := zapcore.NewConsoleEncoder(defaultEncoderConfig)
	level := zap.DebugLevel
	if !debug {
		defaultEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		level = zap.InfoLevel
		encoder = zapcore.NewJSONEncoder(defaultEncoderConfig)
	}

	instance = zap.New(zapcore.NewCore(encoder, output, level), zap.AddCaller()).Sugar().Named(appName)
	initPrinter(instance)
}

// Close log sync and remove instance
func Close() {
	_ = instance.Sync()
	instance = nil
}

// ShowLine Line number display switch
func ShowLine(enable bool) *zap.SugaredLogger {
	return instance.Desugar().WithOptions(zap.WithCaller(enable)).Sugar()
}

// Output return a log io.Writer
func Output() io.Writer {
	return output
}
