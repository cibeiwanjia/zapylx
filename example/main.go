package main

import (
	"github.com/cibeiwanjia/zapylx/logger"
	"go.uber.org/zap"
	"time"
)

func main() {
	// 示例 1: 使用默认配置
	logger.InitWithDefault()

	logger.Info("使用默认配置启动")
	logger.Debug("这是调试信息")
	logger.Warn("这是警告信息")
	logger.Errorf("这是错误信息: %s", "测试错误")

	logger.Sync()

	// 示例 2: 使用自定义配置
	config := &logger.Config{
		Level:         "debug",
		Filename:      "./logs/example.log",
		MaxSize:       10,
		MaxBackups:    5,
		MaxAge:        30,
		Compress:      true,
		ConsoleOutput: true,
		JSONOutput:    false,
		CallerSkip:    1,
	}

	logger.InitLogger(config)

	// 记录不同级别的日志
	logger.Debug("Debug 模式已启用")
	logger.Info("应用程序启动成功", zap.String("version", "1.0.0"))
	logger.Warn("配置文件未找到，使用默认配置")

	// 记录带有错误堆栈的错误日志
	logger.Error("处理请求失败",
		zap.Int("statusCode", 500),
		zap.Duration("duration", 1234*time.Millisecond),
	)

	// 记录带有详细信息的日志
	logger.Info("用户登录",
		zap.String("username", "alice"),
		zap.Int("userID", 12345),
		zap.Time("loginTime", time.Now()),
	)

	// 使用 With 创建带有固定字段的子 logger
	requestLogger := logger.With(
		zap.String("requestID", "abc-123"),
		zap.String("ip", "192.168.1.1"),
	)

	requestLogger.Info("处理请求")
	requestLogger.Info("请求处理完成", zap.Duration("elapsed", 567*time.Millisecond))

	logger.Sync()
}
