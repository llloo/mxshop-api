package main

import (
	"fmt"
	"go.uber.org/zap"
	"mxshop-api/user-api/initialize"
)

func main() {
	port := 8000

	initialize.InitLogger()

	r := initialize.InitRouters()
	zap.S().Infof("启动服务， port: %d", port)
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panicf("启动服务失败， port: %d", port)
	}
}
