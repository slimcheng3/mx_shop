package main

import (
	"fmt"
	"mxshop_api/user-web/initialize"

	"go.uber.org/zap"
)

func main()  {
	// 1、初始化日志
	initialize.InitLogger()

	// 2、初始化路由
	router := initialize.InitRouter()
	zap.S().Info("################ start gin server ####################")
	if err := router.Run(fmt.Sprintf(":%d", 8585)); err != nil{
		zap.S().Panic("启动失败：", err.Error())
	}
}