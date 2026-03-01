package main

import (
	"github.com/cibeiwanjia/zapylx"
	"go.uber.org/zap"
	"time"
)

func main() {
	// 使用默认配置初始化日志
	logger.InitWithDefault()

	logger.Info("使用默认配置启动")
	logger.Debugf("当前时间: %s", time.Now().Format("2006-01-02 15:04:05"))
	logger.Warn("这是警告信息")

	// 记录带有字段的日志
	logger.Info("记录详细信息",
		zap.String("service", "example"),
		zap.Int("count", 100),
	)

	logger.Errorf("这是错误信息: %s", "测试错误")

	logger.Sync()
}
