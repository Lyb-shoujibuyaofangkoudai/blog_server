package flag

import (
	"blog_server/global"
	sys_flag "flag"
)

type Option struct {
	DB bool
}

// Parse 解析命令行： go run main.go -db
func Parse() Option {
	// go run main.go -db
	db := sys_flag.Bool("db", false, "初始化数据库")
	// 解析命令
	sys_flag.Parse()
	return Option{
		DB: *db,
	}
}

func IsStopWeb(option *Option) bool {
	if option.DB {
		global.Log.Infof("停止web项目")
		return true
	}
	return false
}

func SwitchOption(option *Option) {
	if option.DB {
		// 迁移数据库
		MakeMigration()
	}
}
