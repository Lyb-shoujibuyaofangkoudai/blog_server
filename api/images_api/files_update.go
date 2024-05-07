package images_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type FileUpdate struct {
	ID   uint   `json:"id" binding:"required" msg:"id参数不能为空"`
	Name string `json:"name" binding:"required" msg:"name参数不能为空"`
}

// FileUpdateView 修改文件名称
func (ImagesApi) FileUpdateView(c *gin.Context) {
	cr := FileUpdate{}
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithValidateError(err, &cr, c)
		return
	}
	global.Log.Infof("fileTypeParams: %v", cr)
	rowsAffected := global.DB.Find(&models.ImageModel{
		MODEL: models.MODEL{ID: cr.ID},
	}).RowsAffected
	global.Log.Infof("rowsAffected: %v", rowsAffected)
	if rowsAffected == 0 {
		res.FailWithMsg("文件不存在", c)
		return
	}
	err = global.DB.Model(&models.ImageModel{}).Where("id = ?", cr.ID).Update("name", cr.Name).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改失败", c)
		return
	}
	res.OkWith(c)
}
