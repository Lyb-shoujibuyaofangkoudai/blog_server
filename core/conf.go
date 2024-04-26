package core

import (
	"blog_server/config"
	"blog_server/global"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const ConfigFilePath = "settings.yaml"

// InitConfig 读取yaml配置
func InitConfig() {
	dataBytes, err := os.ReadFile(ConfigFilePath)
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

func SetYaml() error {
	// 将结果提转换为字节
	// 注意：这里是将整个global.Config 转换为yaml格式的字节，因为后续写入的数据会直接覆盖整个yaml文件，global.Config是指针类型，之前修改的siteInfo也会被保存
	bytes, err := yaml.Marshal(global.Config)
	if err != nil {
		global.Log.Error("yaml 转换失败：", err)
		return err
	}
	err = os.WriteFile(ConfigFilePath, bytes, 0644)
	if err != nil {
		global.Log.Errorf("修改%s文件失败：%v", ConfigFilePath, err)
		return err
	}

	global.Log.Infof("修改%s文件成功", ConfigFilePath)
	return nil
}
