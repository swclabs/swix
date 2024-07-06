package main

import (
	"fmt"
	"swclabs/swipecore/pkg/lib/logger"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.DateOnly))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.InfoLevel:
		enc.AppendString(fmt.Sprintf("[%s] %s", logger.Green.Add("SWIPE"), logger.Blue.Add(level.CapitalString())))
		return
	case zapcore.DebugLevel:
		enc.AppendString(fmt.Sprintf("[%s] %s", logger.Green.Add("SWIPE"), logger.Magenta.Add(level.CapitalString())))
		return
	case zapcore.WarnLevel:
		enc.AppendString(fmt.Sprintf("[%s] %s", logger.Green.Add("SWIPE"), logger.Yellow.Add(level.CapitalString())))
		return
	}
	enc.AppendString(fmt.Sprintf("[%s] %s", logger.Green.Add("SWIPE"), logger.Red.Add(level.CapitalString())))
}

func main() {
	// Tạo encoder config cho console
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    CustomLevelEncoder,
		EncodeTime:     SyslogTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		// EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Tạo cấu hình cho Zap
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	// Tạo logger từ cấu hình
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// Sử dụng logger
	logger.Info("Thông điệp thông tin")

	// Ghi log các mức khác nhau
	logger.Debug("Thông điệp debug")
	logger.Warn("Thông điệp cảnh báo")
	// logger.Error("Thông điệp lỗi")
}
