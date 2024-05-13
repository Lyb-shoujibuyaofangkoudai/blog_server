package routes

import "blog_server/api"

func (router RouterGroup) MenuRouter() {
	menuApi := api.ApiGroupApp.MenuApi
	router.GET("menu_list_detail", menuApi.MenuListDetailView)
	router.GET("menu_list_sample", menuApi.MenuNameListView)
	router.GET("menu/:id", menuApi.MenuDetailView)
	router.POST("menu", menuApi.MenuAddView)
	router.PUT("menu", menuApi.MenuListUpdate)
	router.DELETE("menu", menuApi.MenuRemoveView)
}
