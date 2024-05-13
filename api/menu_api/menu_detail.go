package menu_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuDetailView(c *gin.Context) {
	id := c.Param("id")
	var menu models.MenuModel
	err := global.DB.Take(&menu, id).Error
	if err != nil {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	res.OkWithData(menu, c)
}
