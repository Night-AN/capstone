package log

import (
	// 你的配置包
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once sync.Once
	log  *zap.Logger
)

func New(logLevel int, logFormat string) *zap.Logger {
	once.Do(func() {
		level := zapcore.Level(logLevel)
		console := zapcore.Lock(zapcore.AddSync(os.Stdout))

		encCfg := zapcore.EncoderConfig{
			MessageKey:       "msg",
			LevelKey:         "level",
			TimeKey:          "ts",
			CallerKey:        "caller",
			EncodeLevel:      zapcore.CapitalColorLevelEncoder,
			EncodeTime:       zapcore.ISO8601TimeEncoder,
			EncodeCaller:     zapcore.ShortCallerEncoder,
			ConsoleSeparator: " ",
		}
		var enc zapcore.Encoder
		if logFormat == "json" {
			enc = zapcore.NewJSONEncoder(encCfg)
		} else {
			enc = zapcore.NewConsoleEncoder(encCfg)
		}

		core := zapcore.NewCore(enc, console, zap.NewAtomicLevelAt(level))
		log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	})
	return log
}

func Get() *zap.Logger {
	return log
}
