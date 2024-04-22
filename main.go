package main

import (
	"blog_server/core"
	"blog_server/global"
	"fmt"
)

func main() {
	core.InitConfig()
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}
