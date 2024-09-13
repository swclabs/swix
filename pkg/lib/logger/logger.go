package logger

import (
	"fmt"
	"io"
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	encoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    customLevelEncoder,
		EncodeTime:     syslogTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		// EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	config = zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
)

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.DateOnly))
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.InfoLevel:
		enc.AppendString(fmt.Sprintf("[%s] %s", Green.Add("SWIPE"), Blue.Add(level.CapitalString())))
		return
	case zapcore.DebugLevel:
		enc.AppendString(fmt.Sprintf("[%s] %s", Green.Add("SWIPE"), Magenta.Add(level.CapitalString())))
		return
	case zapcore.WarnLevel:
		enc.AppendString(fmt.Sprintf("[%s] %s", Green.Add("SWIPE"), Yellow.Add(level.CapitalString())))
		return
	}
	enc.AppendString(fmt.Sprintf("[%s] %s", Green.Add("SWIPE"), Red.Add(level.CapitalString())))
}

// Info logs an info message
func Info(msg string) {
	logger, err := config.Build()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			return // ignore error
		}
	}()
	logger.Info(msg)
}

// Debug logs an Debug message
func Debug(msg string) {
	logger, err := config.Build()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			return // ignore error
		}
	}()
	logger.Debug(msg)
}

// Write writes logs to a file
func Write(file io.Writer) *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	writer := zapcore.AddSync(file)
	core := zapcore.NewCore(encoder, writer, zap.InfoLevel)
	return zap.New(core)
}
