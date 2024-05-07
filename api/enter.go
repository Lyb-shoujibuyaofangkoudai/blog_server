package api

import (
	"blog_server/api/adverts_api"
	"blog_server/api/images_api"
	"blog_server/api/settings_api"
)

// ApiGroup 统一导出的api
type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertsApi  adverts_api.AdvertApi
}

// ApiGroupApp new函数实例化，实例化完成后会返回结构体地指针类型
var ApiGroupApp = new(ApiGroup)
