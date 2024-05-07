package routes

import "blog_server/api"

func (router RouterGroup) ImagesRoutes() {
	imagesApi := api.ApiGrounpApp.ImagesApi
	router.POST("file/", imagesApi.FileUploadView)
	router.POST("files/", imagesApi.FilesUploadViews)
	router.GET("files/:type", imagesApi.FilesListView)
	router.DELETE("files", imagesApi.FileRemoveView)
	router.PUT("file", imagesApi.FileUpdateView)
}
