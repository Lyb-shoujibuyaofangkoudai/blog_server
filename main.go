package main

import (
	"blog_server/core"
	"blog_server/global"
)

func main() {
	// 初始化配置文件（读取配置文件）
	core.InitConfig()
	// 初始化啊日志
	global.Log = core.InitLogger()
	global.Log.Warnf("初始化数据库")
	global.Log.Error("错误")
	global.Log.Infof("信息")
	// 初始化数据库
	global.DB = core.InitGorm()
}
