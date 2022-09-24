package main

import (
	"github.com/RockRockWhite/minio-client/config"
	"github.com/RockRockWhite/minio-client/routers"
	"github.com/RockRockWhite/minio-client/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// 初始化配置文件
	config.Init("./config/config.yml")

	// 初始化logger
	utils.InitLogger(viper.GetString("Logger.LogFile"), logrus.DebugLevel, "2006-01-02 15:04:05")
	utils.Logger().Infof("| [service] | ***** Service started ***** |")
	defer utils.Logger().Infof("| [service] | ***** Service stoped ***** |")
	// 初始化Utils
	utils.InitJwt(viper.GetString("Jwt.Secret"), viper.GetString("Jwt.Issuer"), viper.GetInt("Jwt.ExpireDays"))

	// 初始化并运行路由
	router := routers.InitApiRouter()

	_ = router.Run(viper.GetString("HttpServer.Port"))
}
