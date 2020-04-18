package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"strings"
	"time"
)

func NewLogger() *zap.Logger {
	//// 最后创建具体的Logger
	conf := zap.NewProductionEncoderConfig()
	conf.MessageKey = "Flag"
	conf.LevelKey = "Level"
	conf.EncodeLevel = zapcore.CapitalLevelEncoder
	conf.TimeKey = "Ts"
	conf.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	cnf := zapcore.NewJSONEncoder(conf)
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(cnf, zapcore.AddSync(rolling("info.log")), infoLevel),
		zapcore.NewCore(cnf, zapcore.AddSync(rolling("error.log")), errorLevel),
	)
	// 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
	return zap.New(core)
}

func rolling(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   strings.Join([]string{"./runtime/", filename}, ""),
		MaxSize:    100, // megabytes
		MaxBackups: 1,
		MaxAge:     7, // days
		Compress:  true,
	}
}
