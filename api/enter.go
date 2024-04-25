package api

import "blog_server/api/settings_api"

type ApiGrounp struct {
	SettingsApi settings_api.SettingsApi
}

// new函数实例化，实例化完成后会返回结构体地指针类型
var ApiGrounpApp = new(ApiGrounp)
