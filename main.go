package main

import (
	"blog_server/core"
	"blog_server/flag"
	"blog_server/global"
	"blog_server/models/res"
	"blog_server/routes"
	"blog_server/utils/validators"
)

func main() {
	// 初始化配置文件（读取配置文件）
	core.InitConfig()
	// 初始化啊日志
	global.Log = core.InitLogger()
	res.ReadErrorCodeJson()
	// 初始化数据库
	global.DB = core.InitGorm()
	// 初始化Redis
	global.Redis = core.InitRedis()

	option := flag.Parse()
	if flag.IsStopWeb(&option) {
		flag.SwitchOption(&option)
		return
	}

	router := routes.InitRoutes()
	// 注册自定义校验器
	validators.RegisterPhoneValidators()
	validators.LoginCodeValidate()

	addr := global.Config.System.Addr()
	global.Log.Infof("程序运行在：%s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf("程序启动失败：%s", err.Error())
	}
}
