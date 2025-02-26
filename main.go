package main

import (
	"backend/config"
	"backend/log"
	"backend/logo"
	"backend/router"
)

func main() {
	// 初始化配置文件
	config.MustInit()
	// 初始化日志
	log.Setup(&config.Get().Log)

	// 打印logo
	logo.Print()

	// 初始化路由
	router.Setup()
}
