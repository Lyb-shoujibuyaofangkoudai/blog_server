package main

import (
	"blog_server/core"
	"blog_server/global"
	"blog_server/routes"
)

func main() {
	// 初始化配置文件（读取配置文件）
	core.InitConfig()
	// 初始化啊日志
	global.Log = core.InitLogger()
	// 初始化数据库
	global.DB = core.InitGorm()

	router := routes.InitRoutes()
	addr := global.Config.System.Addr()
	global.Log.Infof("程序运行在：%s", addr)
	router.Run(addr)
}
