package routes

import "blog_server/api"

func (router RouterGroup) AdvertRouter() {
	advertsApi := api.ApiGroupApp.AdvertsApi
	router.GET("/adverts", advertsApi.AdvertList)
	router.POST("/adverts", advertsApi.AdvertAdd)
	router.PUT("/adverts/:advertId", advertsApi.AdvertUpdateView)
	router.DELETE("/adverts", advertsApi.AdvertRemoveView)
}
