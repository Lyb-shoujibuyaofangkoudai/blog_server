package settings_api

import (
	"blog_server/config"
	"blog_server/core"
	"blog_server/global"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingInfoUpdateView(c *gin.Context) {
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}
	// 注意：这里修改的是引用yaml的内容（global.Config.SiteInfo），不会修改yaml源文件中的内容
	// 注意：这里的global.Config.SiteInfo是一个指针，是一个引用所以可以修改原数据
	// 注意：会直接覆盖整个SiteInfo，所以传过来的数据要全
	global.Config.SiteInfo = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error("修改配置文件中的站点数据失败", err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OkWith(c)

}
