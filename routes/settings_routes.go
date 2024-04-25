package routes

import (
	"blog_server/api"
)

func (router RouterGroup) SettingsRoutes() {
	settingsApi := api.ApiGrounpApp.SettingsApi
	router.GET("settings", settingsApi.SettingInfoView)
}
