package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	once   sync.Once
	Logger *zap.Logger
)

// Config 日志配置
type Config struct {
	// Level 日志级别: debug, info, warn, error, fatal
	Level string
	// Filename 日志文件路径，为空则不输出到文件
	Filename string
	// MaxSize 单个日志文件最大大小(MB)
	MaxSize int
	// MaxBackups 保留的旧日志文件最大数量
	MaxBackups int
	// MaxAge 保留旧日志文件的最大天数
	MaxAge int
	// Compress 是否压缩旧日志文件
	Compress bool
	// ConsoleOutput 是否输出到控制台
	ConsoleOutput bool
	// JSONOutput 是否使用JSON格式输出控制台日志
	JSONOutput bool
	// CallerSkip 调用栈跳过层数
	CallerSkip int
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Level:         "info",
		Filename:      "./logs/app.log",
		MaxSize:       100,
		MaxBackups:    3,
		MaxAge:        7,
		Compress:      true,
		ConsoleOutput: true,
		JSONOutput:    false,
		CallerSkip:    1,
	}
}

// InitLogger 初始化日志
// 支持日志级别控制、日志切割、JSON格式输出
func InitLogger(config *Config) {
	if config == nil {
		config = DefaultConfig()
	}

	once.Do(func() {
		level := getLogLevel(config.Level)

		var cores []zapcore.Core

		// 文件输出
		if config.Filename != "" {
			hook := lumberjack.Logger{
				Filename:   config.Filename,
				MaxSize:    config.MaxSize,
				MaxBackups: config.MaxBackups,
				MaxAge:     config.MaxAge,
				Compress:   config.Compress,
				LocalTime:  true,
			}

			encoderConfig := getEncoderConfig()
			core := zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),
				zapcore.AddSync(&hook),
				level,
			)
			cores = append(cores, core)
		}

		// 控制台输出
		if config.ConsoleOutput {
			encoderConfig := getEncoderConfig()
			var encoder zapcore.Encoder
			if config.JSONOutput {
				encoder = zapcore.NewJSONEncoder(encoderConfig)
			} else {
				encoder = zapcore.NewConsoleEncoder(encoderConfig)
			}
			core := zapcore.NewCore(
				encoder,
				zapcore.AddSync(os.Stdout),
				level,
			)
			cores = append(cores, core)
		}

		var core zapcore.Core
		if len(cores) > 0 {
			core = zapcore.NewTee(cores...)
		} else {
			// 默认输出到控制台
			encoderConfig := getEncoderConfig()
			core = zapcore.NewCore(
				zapcore.NewConsoleEncoder(encoderConfig),
				zapcore.AddSync(os.Stdout),
				level,
			)
		}

		Logger = zap.New(core,
			zap.AddCaller(),
			zap.AddCallerSkip(config.CallerSkip),
			zap.AddStacktrace(zapcore.ErrorLevel),
		)
	})
}

// getEncoderConfig 返回编码器配置
func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// getLogLevel 将字符串级别转换为 zapcore.Level
func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

// Sync 同步日志缓冲区
func Sync() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}

// GetLogger 返回底层 logger 实例
func GetLogger() *zap.Logger {
	return Logger
}
