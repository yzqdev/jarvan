package main

import (
	"fmt"
	"jarvan/src/models"
	"jarvan/src/pkg/gredis"
	"jarvan/src/pkg/logging"
	"jarvan/src/pkg/setting"
	"jarvan/src/routers"

	"github.com/fvbock/endless"
	"github.com/labstack/gommon/log"
	"syscall"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()

	// ------------------- use endless 热更新，不需要重启服务 ----------------//
	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server error: %v", err)
	}
}
