package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogLevel string

type initLoggerOption func(*zap.Config)

const (
	LogLevelDebug   LogLevel = "debug"
	LogLevelInfo    LogLevel = "info"
	LogLevelWarning LogLevel = "warning"
	LogLevelError   LogLevel = "error"
	LogLevelDPanic  LogLevel = "dpanic"
	LogLevelPanic   LogLevel = "panic"
	LogLevelFatal   LogLevel = "fatal"
)

const (
	jsonEncoding      = "json"
	defaultOutputSink = "stdout"
	defaultErrorSink  = "stderr"
	levelKey          = "level"
	timeKey           = "timestamp"
	nameKey           = "logger"
	callerKey         = "caller"
	stacktraceKey     = "stacktrace"
	messageKey        = "msg"
)

var (
	_logger *zap.Logger
)

func InitLogger(options ...initLoggerOption) error {
	config := buildLoggerConfig()

	for _, applyOption := range options {
		applyOption(config)
	}

	logger, err := config.Build(zap.AddStacktrace(zapcore.DPanicLevel))
	if err != nil {
		return err
	}

	logger.WithOptions()

	_logger = logger

	return nil
}

func Logger() *zap.Logger {
	return _logger
}

func InitLoggerWithLevelOption(level LogLevel) initLoggerOption {
	return func(c *zap.Config) {
		c.Level = zap.NewAtomicLevelAt(getZapLevel(level))
	}
}

func InitLoggerWithInitialFieldsOption(initialFields map[string]interface{}) initLoggerOption {
	return func(c *zap.Config) {
		c.InitialFields = initialFields
	}
}

func InitLoggerWithPathOption(paths []string) initLoggerOption {
	return func(c *zap.Config) {
		var filePaths []string
		for _, dir := range paths {
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				os.Mkdir(dir, os.ModePerm)
			}

			fmt.Println(dir)
			filePath := fmt.Sprintf("%s/service_%d.log", dir, time.Now().UnixNano())
			_, err := os.Create(filePath)
			if err != nil {
				fmt.Println("could not create file. ", filePath)
				continue
			}

			filePaths = append(filePaths, filePath)
		}

		if len(filePaths) > 0 {
			c.OutputPaths = filePaths
		}
	}
}

func getZapLevel(level LogLevel) zapcore.Level {
	switch level {
	case LogLevelDebug:
		return zapcore.DebugLevel
	case LogLevelInfo:
		return zapcore.InfoLevel
	case LogLevelWarning:
		return zapcore.WarnLevel
	case LogLevelError:
		return zapcore.ErrorLevel
	case LogLevelDPanic:
		return zapcore.DPanicLevel
	case LogLevelPanic:
		return zapcore.PanicLevel
	case LogLevelFatal:
		return zapcore.FatalLevel
	default:
		return zapcore.WarnLevel
	}
}

func buildLoggerConfig() *zap.Config {
	config := zap.Config{
		Encoding:         jsonEncoding,
		OutputPaths:      []string{defaultOutputSink},
		ErrorOutputPaths: []string{defaultErrorSink},
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:       levelKey,
			TimeKey:        timeKey,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			NameKey:        nameKey,
			CallerKey:      callerKey,
			StacktraceKey:  stacktraceKey,
			MessageKey:     messageKey,
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	return &config
}
