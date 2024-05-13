package menu_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// MenuRemoveView 删除菜单
func (MenuApi) MenuRemoveView(c *gin.Context) {
	ids := common.RemoveFileList{}
	if err := c.ShouldBindJSON(&ids); err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	global.Log.Infof("ids: %v", ids.Ids)
	var menuModels []models.MenuModel
	count := global.DB.Find(&menuModels, ids.Ids).RowsAffected
	if count == 0 {
		res.FailWithMsg("没有找到该菜单", c)
		return
	}

	// 开启事务，如果返回错误，事务会回滚
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		err := global.DB.Model(&menuModels).Association("MenuImages").Clear()
		if err != nil {
			return err
		}
		err = global.DB.Delete(&menuModels).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OkWithMsg(fmt.Sprintf("删除了%d个菜单", count), c)
}
