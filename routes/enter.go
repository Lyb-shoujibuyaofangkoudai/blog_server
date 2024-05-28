package routes

import (
	"blog_server/global"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRoutes() *gin.Engine {

	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.Use(cors.Default()) // 解决跨域问题
	routerGroup := router.Group("api")
	routerGroupApp := RouterGroup{routerGroup}

	// 系统配置API
	routerGroupApp.SettingsRoutes()
	routerGroupApp.ImagesRoutes()
	routerGroupApp.AdvertRouter()
	routerGroupApp.MenuRouter()
	routerGroupApp.UserRouter()
	routerGroupApp.CodeRouter()
	return router
}
