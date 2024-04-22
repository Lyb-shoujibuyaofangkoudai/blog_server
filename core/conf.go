package core

import (
	"blog_server/config"
	"blog_server/global"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// 读取yarm配置
func InitConfig() {
	dataBytes, err := os.ReadFile("settings.yaml")
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	//fmt.Println("yaml 文件的内容: \n", string(dataBytes))
	c := config.Config{}
	err = yaml.Unmarshal(dataBytes, &c)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return
	}
	//fmt.Printf("c → %+v\n", c) // c → {Mysql:{Url:127.0.0.1 Port:3306} Redis:{Host:127.0.0.1 Port:6379}}
	global.Config = &c // 存入全局变量
}
