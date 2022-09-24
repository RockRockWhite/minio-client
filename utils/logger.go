package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger

// InitLogger 初始化logger
func InitLogger(logFile string, level logrus.Level, timestampFormat string) {
	// 配置log
	file, _ := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	logger = logrus.New()
	logger.Out = file
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: timestampFormat})
}

// Logger 获得logger
func Logger() *logrus.Logger {
	return logger
}
