package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	// zap.NewDevelopment 格式化输出
	// zap.NewProduction json序列化输出
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	logger.Info("无法获取网址",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}