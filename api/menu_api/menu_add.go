package menu_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type MenuRequest struct {
	Id            uint               `json:"id"`
	MenuTitle     string             `json:"menu_title" binding:"required" msg:"menu_title参数不能为空"`
	MenuTitleEn   string             `json:"menu_title_en"`
	Path          string             `json:"path" binding:"required" msg:"path参数不能为空"`
	Slogan        string             `json:"slogan"`
	Abstract      ctype.Array        `json:"abstract" binding:"required" msg:"abstract参数不能为空"`
	AbstractTime  int                `json:"abstract_time" binding:"required" msg:"abstract_time参数不能为空"`
	MenuTime      int                `json:"menu_time"`
	Sort          int                `json:"sort" binding:"required"`
	Icon          string             `json:"icon"`
	ImageSortList []models.ImageSort `json:"image_sort_list"`
}

func (MenuApi) MenuAddView(c *gin.Context) {
	var menuRequest MenuRequest
	err := c.ShouldBindJSON(&menuRequest)
	if err != nil {
		res.FailWithMsg(fmt.Sprintf("参数有误: %v", err), c)
		return
	}
	global.Log.Infof("menuRequest: %v", menuRequest)

	// 校验菜单标题和菜单路径是否已经存在
	count := global.DB.Where("menu_title = ?", menuRequest.MenuTitle).Or("path = ?", menuRequest.Path).First(&models.MenuModel{}).RowsAffected
	if count > 0 {
		res.FailWithMsg("菜单标题或菜单路径已存在", c)
		return
	}
	menuModel := models.MenuModel{
		MenuTitle:    menuRequest.MenuTitle,
		MenuTitleEn:  menuRequest.MenuTitleEn,
		Slogan:       menuRequest.Slogan,
		Abstract:     menuRequest.Abstract,
		AbstractTime: menuRequest.AbstractTime,
		MenuTime:     menuRequest.MenuTime,
		Sort:         menuRequest.Sort,
		Path:         menuRequest.Path,
		Icon:         menuRequest.Icon,
	}
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(fmt.Sprintf("MenuModel添加失败：%v", err))
		res.FailWithMsg("添加失败", c)
		return
	}
	var menuImageModels []models.MenuImageModel
	// 判断是否有图片ImageSort
	if len(menuRequest.ImageSortList) > 0 {
		for _, v := range menuRequest.ImageSortList {
			menuImageModel := models.MenuImageModel{
				MenuID:  menuModel.ID,
				Sort:    v.Sort,
				ImageID: v.ImageID,
			}
			menuImageModels = append(menuImageModels, menuImageModel)
		}
	}
	err = global.DB.Create(&menuImageModels).Error
	if err != nil {
		global.Log.Error(fmt.Sprintf("MenuImageModel添加失败：%v", err))
		res.FailWithMsg("关联图片失败", c)
		return
	}

	res.OkWithMsg("添加成功", c)
}
