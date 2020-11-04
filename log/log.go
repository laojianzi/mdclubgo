package log

import (
	"io"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/laojianzi/mdclubgo/conf"
)

var instance *zap.SugaredLogger

var output zapcore.WriteSyncer

// Init log instance
func Init() {
	ws := []zapcore.WriteSyncer{zapcore.AddSync(os.Stdout)}

	if logPath := conf.Log.RootPath; logPath != "" {
		if logPath[0] != '/' {
			logPath = filepath.Join(conf.WorkDir(), logPath)
		}

		ws = append(ws, zapcore.AddSync(&lumberjack.Logger{
			Filename:   logPath,
			MaxSize:    500,
			MaxBackups: 3,
			MaxAge:     28,
			Compress:   true,
		}))
	}

	output = zapcore.NewMultiWriteSyncer(ws...)
	encoderLevel := zapcore.CapitalColorLevelEncoder
	if !conf.App.Debug {
		encoderLevel = zapcore.CapitalLevelEncoder
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "TIME",
		LevelKey:       "LEVEL",
		NameKey:        "LOGGER",
		CallerKey:      "LINE",
		MessageKey:     "MESSAGE",
		StacktraceKey:  "STACKTRACE",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encoderLevel,               // 大写编码器，debug 模式时带颜色
		EncodeTime:     zapcore.ISO8601TimeEncoder, // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	var encoder zapcore.Encoder
	var level zapcore.Level
	if conf.App.Debug {
		level = zap.DebugLevel
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		level = zap.InfoLevel
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	instance = zap.New(zapcore.NewCore(encoder, output, level), zap.AddCaller()).Sugar().Named(conf.App.Name)
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
