package settings_api

import (
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingInfoView(c *gin.Context) {
	res.FailWithCode(res.ParamsError, c)
}
