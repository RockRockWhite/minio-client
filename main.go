package main

import (
	"github.com/RockRockWhite/minio-client/config"
	"github.com/RockRockWhite/minio-client/routers"
)

func main() {
	router := routers.InitApiRouter()
	_ = router.Run(config.GetString("HttpServer.Port"))
}
