package settings_api

import "github.com/gin-gonic/gin"

func (SettingsApi) SettingInfoView(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "哈哈哈123123",
	})
}
