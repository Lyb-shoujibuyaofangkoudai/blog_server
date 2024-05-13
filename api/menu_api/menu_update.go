package menu_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

// MenuListUpdate 更新菜单
func (MenuApi) MenuListUpdate(c *gin.Context) {
	var menuRequest MenuRequest
	err := c.ShouldBindJSON(&menuRequest)
	if err != nil {
		res.FailWithMsg(fmt.Sprintf("参数有误: %v", err), c)
		return
	}
	if menuRequest.Id == 0 {
		res.FailWithMsg("id不能为空", c)
		return
	}
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, menuRequest.Id).Error
	if err != nil {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	menuModel.ID = menuRequest.Id
	menuModel.Abstract = menuRequest.Abstract
	menuModel.AbstractTime = menuRequest.AbstractTime
	menuModel.MenuTitle = menuRequest.MenuTitle
	menuModel.MenuTitleEn = menuRequest.MenuTitleEn
	menuModel.MenuTime = menuRequest.MenuTime
	menuModel.Path = menuRequest.Path
	menuModel.Slogan = menuRequest.Slogan
	menuModel.Sort = menuRequest.Sort
	if menuRequest.Icon != "" {
		menuModel.Icon = menuRequest.Icon
	}

	err = global.DB.Where("id = ?", menuRequest.Id).Updates(&menuModel).Error
	if err != nil {
		res.FailWithMsg("更新失败", c)
		return
	}

	// 是否传入了image_sort_list
	if len(menuRequest.ImageSortList) > 0 {
		// 删除之前的关联
		err = global.DB.Where("menu_id = ?", menuRequest.Id).Delete(&models.MenuImageModel{}).Error
		if err != nil {
			res.FailWithMsg("删除关联失败", c)
			return
		}
		// 关联新的图片数据
		for _, imageSort := range menuRequest.ImageSortList {
			err = global.DB.Create(&models.MenuImageModel{
				MenuID:  menuRequest.Id,
				ImageID: imageSort.ImageID,
				Sort:    imageSort.Sort,
			}).Error
			if err != nil {
				res.FailWithMsg("关联失败", c)
				return
			}
		}
	}

	res.OkWithMsg("更新成功", c)
}
