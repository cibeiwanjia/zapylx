package logger

import (
	"go.uber.org/zap"
)

// InitWithDefault 使用默认配置初始化日志
func InitWithDefault() {
	InitLogger(nil)
}

// Info logs a message at InfoLevel
func Info(msg string, fields ...zap.Field) {
	if Logger == nil {
		InitWithDefault()
	}
	Logger.Info(msg, fields...)
}

// Error logs a message at ErrorLevel
func Error(msg string, fields ...zap.Field) {
	if Logger == nil {
		InitWithDefault()
	}
	Logger.Error(msg, fields...)
}

// Debug logs a message at DebugLevel
func Debug(msg string, fields ...zap.Field) {
	if Logger == nil {
		InitWithDefault()
	}
	Logger.Debug(msg, fields...)
}

// Warn logs a message at WarnLevel
func Warn(msg string, fields ...zap.Field) {
	if Logger == nil {
		InitWithDefault()
	}
	Logger.Warn(msg, fields...)
}

// Fatal logs a message at FatalLevel and exits
func Fatal(msg string, fields ...zap.Field) {
	if Logger == nil {
		InitWithDefault()
	}
	Logger.Fatal(msg, fields...)
}

// With creates a child logger and adds structured context to it
func With(fields ...zap.Field) *zap.Logger {
	if Logger == nil {
		InitWithDefault()
	}
	return Logger.With(fields...)
}

// Infof logs a formatted message at InfoLevel
func Infof(format string, args ...interface{}) {
	if Logger == nil {
		InitWithDefault()
	}
	Logger.Sugar().Infof(format, args...)
}

// Errorf logs a formatted message at ErrorLevel
func Errorf(format string, args ...interface{}) {
	if Logger == nil {
		InitWithDefault()
	}
	Logger.Sugar().Errorf(format, args...)
}

// Debugf logs a formatted message at DebugLevel
func Debugf(format string, args ...interface{}) {
	if Logger == nil {
		InitWithDefault()
	}
	Logger.Sugar().Debugf(format, args...)
}

// Warnf logs a formatted message at WarnLevel
func Warnf(format string, args ...interface{}) {
	if Logger == nil {
		InitWithDefault()
	}
	Logger.Sugar().Warnf(format, args...)
}

// Fatalf logs a formatted message at FatalLevel and exits
func Fatalf(format string, args ...interface{}) {
	if Logger == nil {
		InitWithDefault()
	}
	Logger.Sugar().Fatalf(format, args...)
}
