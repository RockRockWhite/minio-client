package utils

import (
	"github.com/RockRockWhite/minio-client/config"
	"io"
	"log"
	"os"
)

var _logger *log.Logger

func init() {
	w := io.MultiWriter()

	f, err := os.OpenFile(config.GetString("Logger.LogFile"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	w = io.MultiWriter(w, f)

	// 判断控制台输出是否开启
	if config.GetString("Logger.Console") == "enable" {
		w = io.MultiWriter(w, os.Stdout)
	}

	_logger = log.New(w, "[minio-client] ", log.LstdFlags|log.Lmsgprefix)
	_logger.Println("======MINIO-CLIENT START======")
}

func GetLogger() *log.Logger {
	return _logger
}
