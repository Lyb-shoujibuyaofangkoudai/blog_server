package menu_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

// MenuListDetailView 获取menu_models表的详细列表数据，包括其关联表数据
func (MenuApi) MenuListDetailView(c *gin.Context) {
	var menuModel []models.MenuModel
	var ids []uint
	global.DB.Find(&menuModel).Select("id").Scan(&ids) // 会将找到的id存入ids中

	//	 查询连表
	var menuImageModel []models.MenuImageModel
	// Preload 预加载会在查询menuImageModel表时，自动将menuImageModel表中的image_id与image_model表中的id进行匹配
	global.DB.Preload("ImageModel").Order("sort desc").Find(&menuImageModel, "menu_id in ?", ids)
	var menuResponse []MenuResponse
	for _, menuModel := range menuModel {
		//var banners []Banner // 这里如果banners没有值的话会返回数据时返回一个nil值
		banners := []Banner{} // 这样子如果没有值的话，会返回一个空数组[]
		for _, menuImageModel := range menuImageModel {
			if menuImageModel.MenuID == menuModel.ID {
				banners = append(banners, Banner{
					ID:   menuImageModel.ImageModel.ID,
					Path: menuImageModel.ImageModel.Path,
				})
			}
		}
		menuResponse = append(menuResponse, MenuResponse{
			MenuModel: menuModel,
			Banners:   banners,
		})
	}
	res.OkWithData(menuResponse, c)
}

type MenuSampleResponse struct {
	MenuTitle   string `json:"menu_title"`
	MenuTitleEn string `json:"menu_title_en"`
	Path        string `json:"path"`
	Sort        int    `json:"sort"`
	Icon        string `json:"icon"`
}

// MenuNameListView 获取菜单名称列表
func (MenuApi) MenuNameListView(c *gin.Context) {
	sampleResponse := []MenuSampleResponse{}
	global.DB.Table("menu_models").Select("id", "menu_title", "menu_title_en", "path", "icon").Scan(&sampleResponse)
	res.OkWithData(sampleResponse, c)
}
