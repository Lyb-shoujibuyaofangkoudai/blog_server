package adverts_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	advertId := c.Param("advertId")
	var addParams AdvertAdd
	// 注意：使用ShouldBind需要再结构体tag上绑定form
	err := c.ShouldBind(&addParams)
	if err != nil {
		res.FailWithValidateError(err, &addParams, c)
		return
	}
	advertModel := models.AdvertModel{}
	rowsAffected := global.DB.Take(&advertModel, "id = ?", advertId).RowsAffected
	if rowsAffected == 0 {
		res.FailWithMsg("广告不存在", c)
		return
	}

	advertModel.IsShow = addParams.IsShow
	advertModel.Title = addParams.Title
	advertModel.Href = addParams.Href
	advertModel.Images = addParams.Images
	err = global.DB.Updates(&advertModel).Error
	//err = global.DB.Updates(map[string]any{
	//	"is_show": addParams.IsShow,
	//	"title":   addParams.Title,
	//	"href":    addParams.Href,
	//	"images":  addParams.Images,
	//}).Error
	if err != nil {
		res.FailWithMsg(fmt.Sprintf("修改广告失败: %v", err), c)
		return
	}
	res.OkWithMsg("修改广告成功", c)
}
