package routes

import (
	"blog_server/api"
)

func (router RouterGroup) SettingsRoutes() {
	settingsApi := api.ApiGrounpApp.SettingsApi
	router.GET("settings/:name", settingsApi.SettingsView)
	router.PUT("settings/:name", settingsApi.SettingsUpdateView)
}
