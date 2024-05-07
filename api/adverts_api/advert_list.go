package adverts_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"github.com/gin-gonic/gin"
	"strings"
)

func (AdvertApi) AdvertList(c *gin.Context) {
	pageQuery := models.Page{}
	if err := c.ShouldBindQuery(&pageQuery); err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}
	option := common.Option{
		Page:   pageQuery,
		Select: nil,
	}
	var advertModel models.AdvertModel
	// todo: 这里一直无法获取到 referer，有传过来，但是值为空不知道为什么
	referer := c.GetHeader("Referer")
	global.Log.Infof("referer: %s", referer)
	if strings.Contains(referer, "admin") {
		global.Log.Infof("此请求是从后台管理系统访问的")
	} else {
		// 不是管理后台系统访问的只返回显示已启用的广告
		isShow := true
		advertModel.IsShow = &isShow
	}
	list, count, err := common.ComList(&advertModel, option)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OkWithList(list, count, c)
}
