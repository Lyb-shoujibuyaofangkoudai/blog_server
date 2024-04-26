package settings_api

import (
	"blog_server/global"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingInfoView(c *gin.Context) {
	res.OkWithData(global.Config.SiteInfo, c)
}
