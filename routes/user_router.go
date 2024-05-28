package routes

import "blog_server/api"

func (router RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.POST("/register", userApi.Register)
	router.POST("/login", userApi.Login)
}
