package images_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ImagesApi) FileRemoveView(c *gin.Context) {
	fileIds := common.RemoveFileList{}
	if err := c.ShouldBindJSON(&fileIds); err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	global.Log.Infof("fileIds: %v", fileIds.Ids)
	var fileList []models.ImageModel
	count := global.DB.Find(&fileList, fileIds.Ids).RowsAffected
	global.Log.Infof("count: %v", count)
	if count == 0 {
		res.FailWithMsg("没有找到该文件", c)
		return
	}
	rowsAffected := global.DB.Delete(&fileList).RowsAffected
	res.OkWithMsg(fmt.Sprintf("删除了%d个文件", rowsAffected), c)
}
