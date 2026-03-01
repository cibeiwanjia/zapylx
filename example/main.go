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
	logger.Debug("这是调试信息")
	logger.Warn("这是警告信息")
	logger.Errorf("这是错误信息: %s", "测试错误")

	logger.Sync()
}
