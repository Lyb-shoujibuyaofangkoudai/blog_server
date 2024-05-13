package adverts_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	advertIds := common.RemoveFileList{}
	if err := c.ShouldBindJSON(&advertIds); err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	global.Log.Infof("advertIds: %v", advertIds.Ids)
	var advertList []models.AdvertModel
	// 使用Find查询 如果有条件的话只能用ID查找，如果需要用其他字段来查询，需要使用Where
	count := global.DB.Find(&advertList, advertIds.Ids).RowsAffected
	global.Log.Infof("count: %v", count)
	if count == 0 {
		res.FailWithMsg("没有找到该广告", c)
		return
	}
	rowsAffected := global.DB.Delete(&advertList).RowsAffected
	res.OkWithMsg(fmt.Sprintf("删除了%d个广告", rowsAffected), c)
}
