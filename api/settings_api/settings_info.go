package settings_api

import (
	"blog_server/config"
	"blog_server/core"
	"blog_server/global"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
)

type Settings struct {
	Name string `json:"name" uri:"name" binding:"required"`
}

// 获取配置
func (SettingsApi) SettingsView(c *gin.Context) {
	cr := Settings{}
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}
	settingsMap := global.Config.CanGetSettings()
	if settingsMap[cr.Name] == "" {
		res.FailWithMsg("没有该配置项", c)
		return
	}
	settingsInfo := global.Config.GetSettingByName(settingsMap[cr.Name])
	res.OkWithData(settingsInfo.Interface(), c)
}

type configType struct {
	Typ reflect.Type
	Set func(interface{}) error
}

var ConfigTypes = map[string]configType{
	"System": {
		Typ: reflect.TypeOf(config.System{}),
		Set: func(obj interface{}) error {
			return global.Config.SetValue(global.Config, "System", obj)
		},
	},
	"SiteInfo": {
		Typ: reflect.TypeOf(config.SiteInfo{}),
		Set: func(obj interface{}) error {
			return global.Config.SetValue(global.Config, "SiteInfo", obj)
		},
	},
	"QQ": {
		Typ: reflect.TypeOf(config.QQ{}),
		Set: func(obj interface{}) error {
			return global.Config.SetValue(global.Config, "QQ", obj)
		},
	},
	"Email": {
		Typ: reflect.TypeOf(config.Email{}),
		Set: func(obj interface{}) error {
			return global.Config.SetValue(global.Config, "Email", obj)
		},
	},
	"Jwt": {
		Typ: reflect.TypeOf(config.Jwt{}),
		Set: func(obj interface{}) error {
			return global.Config.SetValue(global.Config, "Jwt", obj)
		},
	},
	"QiNiu": {
		Typ: reflect.TypeOf(config.QiNiu{}),
		Set: func(obj interface{}) error {
			return global.Config.SetValue(global.Config, "QiNiu", obj)
		},
	},
}

// SettingsUpdateView 设置配置
func (SettingsApi) SettingsUpdateView(c *gin.Context) {
	settingCr := Settings{}
	settingsMap := global.Config.CanGetSettings()
	err := c.ShouldBindUri(&settingCr)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}
	global.Log.Infof("settingsMap: %v", settingsMap[settingCr.Name])
	if ct, ok := ConfigTypes[settingsMap[settingCr.Name]]; ok {
		// 根据配置项名称创建对应的结构体 注意这里是返回结构体指针
		cr := reflect.New(ct.Typ).Interface()
		// 绑定JSON数据到结构体
		if err := c.ShouldBindJSON(cr); err != nil {
			// 错误处理: 记录日志或者返回错误信息
			global.Log.Errorf("Error binding JSON for %s: %v\n", settingsMap[settingCr.Name], err)
			res.FailWithMsg("绑定JSON数据失败", c)
			return
		}

		//  在调用Set之前，通过反射获取指针指向的值
		if err := ct.Set(reflect.ValueOf(cr).Elem().Interface()); err != nil {
			// 错误处理: 记录日志或者返回错误信息
			log.Printf("Error setting value for %s: %v\n", settingsMap[settingCr.Name], err)
			res.FailWithMsg("设置配置项值失败", c)
			return
		}
		err = core.SetYaml()
		if err != nil {
			global.Log.Error("修改配置文件中的站点数据失败", err)
			res.FailWithMsg(err.Error(), c)
			return
		}
		res.OkWithData(cr, c)
	} else {
		res.FailWithMsg("没有该配置项", c)
	}

	//switch settingsMap[settingCr.Name] {
	//case "System":
	//	cr := config.System{}
	//	err = c.ShouldBindJSON(&cr)
	//	err = global.Config.SetValue(global.Config, settingsMap[settingCr.Name], cr)
	//case "SiteInfo":
	//	cr := config.SiteInfo{}
	//	err = c.ShouldBindJSON(&cr)
	//	err = global.Config.SetValue(global.Config, settingsMap[settingCr.Name], cr)
	//case "QQ":
	//	cr := config.QQ{}
	//	err = c.ShouldBindJSON(&cr)
	//	err = global.Config.SetValue(global.Config, settingsMap[settingCr.Name], cr)
	//case "Email":
	//	cr := config.Email{}
	//	err = c.ShouldBindJSON(&cr)
	//	err = global.Config.SetValue(global.Config, settingsMap[settingCr.Name], cr)
	//case "Jwt":
	//	cr := config.Jwt{}
	//	err = c.ShouldBindJSON(&cr)
	//	err = global.Config.SetValue(global.Config, settingsMap[settingCr.Name], cr)
	//case "QiNiu":
	//	cr := config.QiNiu{}
	//	err = c.ShouldBindJSON(&cr)
	//	err = global.Config.SetValue(global.Config, settingsMap[settingCr.Name], cr)
	//default:
	//	res.FailWithMsg("没有该配置项", c)
	//}

	//config.CreateAndSetConfig(settingCr.Name,c,settingsMap[settingCr.Name])
	//if err != nil {
	//	res.FailWithMsg("修改配置项失败", c)
	//	return
	//}

	//err = core.SetYaml()
	//if err != nil {
	//	global.Log.Error("修改配置文件中的站点数据失败", err)
	//	res.FailWithMsg(err.Error(), c)
	//	return
	//}
	//res.OkWith(c)
	//res.OkWithData(global.Config, c)
}
