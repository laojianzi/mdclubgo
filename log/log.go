package log

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var instance *zap.SugaredLogger

// Init log instance
func Init(hook ...io.Writer) {
	if len(hook) == 0 {
		hook = append(hook, zapcore.Lock(os.Stdout))
	}

	var ws []zapcore.WriteSyncer
	for _, v := range hook {
		ws = append(ws, zapcore.AddSync(v))
	}

	syncer := zapcore.NewMultiWriteSyncer(ws...)
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "TIME",
		LevelKey:       "LEVEL",
		NameKey:        "LOGGER",
		CallerKey:      "LINE",
		MessageKey:     "MESSAGE",
		StacktraceKey:  "STACKTRACE",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,       // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,   //
		EncodeCaller:   zapcore.ShortCallerEncoder,       // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, syncer, zap.DebugLevel)

	instance = zap.New(core, zap.AddCaller()).Sugar().Named("MDClubGo")
}

// Close log sync and remove instance
func Close() {
	_ = instance.Sync()
	instance = nil
}
