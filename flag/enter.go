package flag

import (
	"blog_server/global"
	sys_flag "flag"
)

type Option struct {
	DB   bool
	User string // -u admin
}

// Parse 解析命令行： go run main.go -db
func Parse() Option {
	// go run main.go -db
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建超级管理员用户")
	// 解析命令
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

func IsStopWeb(option *Option) bool {
	if option.DB {
		global.Log.Infof("停止web项目")
		return true
	}
	if option.User != "" {
		global.Log.Infof("停止web项目")
		return true
	}
	return false // 停止web项目
}

func SwitchOption(option *Option) {
	if option.DB {
		// 迁移数据库
		MakeMigration()
		return
	}
	//fmt.Println("创建用户", option.User, option.User == "")
	if option.User == "admin" || option.User == "user" {
		// 创建用户
		CreateUser(option.User)
		return
	}
	sys_flag.Usage()
}
