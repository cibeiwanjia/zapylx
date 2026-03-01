# zapylx

基于 [uber-go/zap](https://github.com/uber-go/zap) 的高性能日志库封装，支持日志切割、控制台输出、JSON格式输出等功能。

## 功能特性

- ✅ 高性能日志输出
- ✅ 支持多级别日志 (debug, info, warn, error, fatal)
- ✅ 自动日志切割 (按大小、按天数、按数量)
- ✅ 支持控制台输出和文件输出
- ✅ 支持 JSON 格式和文本格式
- ✅ 调用堆栈信息
- ✅ 线程安全

## 安装

```bash
go get github.com/cibeiwanjia/zapylx
```

## 快速开始

### 使用默认配置

```go
package main

import (
    "github.com/cibeiwanjia/zapylx"
)

func main() {
    // 使用默认配置初始化日志
    logger.InitWithDefault()

    // 记录日志
    logger.Info("Hello, zapylx!")
    logger.Error("Something went wrong", logger.GetLogger().String("error", "test error"))

    // 程序退出前同步日志
    defer logger.Sync()
}
```

### 自定义配置

```go
package main

import (
    "github.com/cibeiwanjia/zapylx"
)

func main() {
    // 自定义配置
    config := &logger.Config{
        Level:         "debug",           // 日志级别
        Filename:      "./logs/app.log", // 日志文件路径
        MaxSize:       100,               // 单个日志文件最大大小(MB)
        MaxBackups:    3,                 // 保留的旧日志文件最大数量
        MaxAge:        7,                 // 保留旧日志文件的最大天数
        Compress:      true,              // 是否压缩旧日志文件
        ConsoleOutput: true,              // 是否输出到控制台
        JSONOutput:    false,             // 控制台是否使用JSON格式输出
        CallerSkip:    1,                 // 调用栈跳过层数
    }

    logger.InitLogger(config)

    // 记录日志
    logger.Info("Application started")
    logger.Debugf("Debug mode enabled: %v", true)

    defer logger.Sync()
}
```

### 仅输出到控制台

```go
package main

import (
    "github.com/cibeiwanjia/zapylx"
)

func main() {
    config := &logger.Config{
        Level:         "info",
        Filename:      "", // 空字符串表示不输出到文件
        ConsoleOutput: true,
    }

    logger.InitLogger(config)
    defer logger.Sync()
}
```

## API 文档

### 初始化函数

| 函数 | 说明 |
|------|------|
| `InitWithDefault()` | 使用默认配置初始化日志 |
| `InitLogger(config *Config)` | 使用自定义配置初始化日志 |

### 日志记录函数

| 函数 | 说明 |
|------|------|
| `Info(msg string, fields ...zap.Field)` | 记录 Info 级别日志 |
| `Error(msg string, fields ...zap.Field)` | 记录 Error 级别日志 |
| `Debug(msg string, fields ...zap.Field)` | 记录 Debug 级别日志 |
| `Warn(msg string, fields ...zap.Field)` | 记录 Warn 级别日志 |
| `Fatal(msg string, fields ...zap.Field)` | 记录 Fatal 级别日志并退出程序 |
| `Infof(format string, args ...interface{})` | 记录格式化 Info 日志 |
| `Errorf(format string, args ...interface{})` | 记录格式化 Error 日志 |
| `Debugf(format string, args ...interface{})` | 记录格式化 Debug 日志 |
| `Warnf(format string, args ...interface{})` | 记录格式化 Warn 日志 |
| `Fatalf(format string, args ...interface{})` | 记录格式化 Fatal 日志并退出程序 |

### 辅助函数

| 函数 | 说明 |
|------|------|
| `With(fields ...zap.Field) *zap.Logger` | 创建带有固定字段的子 logger |
| `Sync()` | 同步日志缓冲区 |
| `GetLogger() *zap.Logger` | 获取底层 logger 实例 |

## 使用示例

### 结构化日志

```go
logger.Info("User login",
    zap.String("username", "alice"),
    zap.Int("userID", 12345),
    zap.Duration("duration", 1234*time.Millisecond),
)
```

### 格式化日志

```go
logger.Infof("User %s logged in successfully", "alice")
logger.Errorf("Failed to connect to %s:%d", "localhost", 8080)
```

### 使用 With 创建子 logger

```go
requestLogger := logger.With(
    zap.String("requestID", "12345"),
    zap.String("ip", "192.168.1.1"),
)

requestLogger.Info("Processing request")
requestLogger.Error("Request failed", zap.Int("statusCode", 500))
```

### 自定义字段类型

```go
logger.Info("Processing data",
    zap.String("name", "test"),
    zap.Int("count", 100),
    zap.Bool("enabled", true),
    zap.Any("data", map[string]interface{}{"key": "value"}),
)
```

## 配置说明

### Config 结构体

| 字段 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| Level | string | "info" | 日志级别: debug, info, warn, error, fatal |
| Filename | string | "./logs/app.log" | 日志文件路径，为空则不输出到文件 |
| MaxSize | int | 100 | 单个日志文件最大大小(MB) |
| MaxBackups | int | 3 | 保留的旧日志文件最大数量 |
| MaxAge | int | 7 | 保留旧日志文件的最大天数 |
| Compress | bool | true | 是否压缩旧日志文件 |
| ConsoleOutput | bool | true | 是否输出到控制台 |
| JSONOutput | bool | false | 控制台是否使用JSON格式输出 |
| CallerSkip | int | 1 | 调用栈跳过层数 |

## License

MIT
