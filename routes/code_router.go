package routes

import "blog_server/api"

func (router RouterGroup) CodeRouter() {
	codeApi := api.ApiGroupApp.CodeApi
	router.GET("/code/:type", codeApi.Code)
}
