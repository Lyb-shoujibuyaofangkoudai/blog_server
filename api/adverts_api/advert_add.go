package adverts_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AdvertAdd struct {
	Title  string `form:"title" binding:"required" msg:"title参数不能为空"`
	Href   string `form:"href" binding:"required,url" msg:"href参数不能为空或url格式错误"` //url会自动校验路径合法性
	Images string `form:"images" binding:"required,url" msg:"images参数不能为空或url格式错误"`
	IsShow *bool  `form:"is_show"`
}

func (AdvertApi) AdvertAdd(c *gin.Context) {
	var addParams AdvertAdd
	// 注意：使用ShouldBind需要再结构体tag上绑定form
	err := c.ShouldBind(&addParams)
	if err != nil {
		res.FailWithValidateError(err, &addParams, c)
		return
	}
	data := models.AdvertModel{
		Title: addParams.Title,
	}
	global.Log.Infof("查询到的数据：%v", data)
	rowsAffected := global.DB.Where(&data).Find(&data).RowsAffected
	if rowsAffected > 0 {
		res.FailWithMsg(fmt.Sprintf("标题为：%s 的广告已存在", addParams.Title), c)
		return
	}

	data.Href = addParams.Href
	data.Images = addParams.Images
	data.IsShow = addParams.IsShow
	err = global.DB.Create(&data).Error
	if err != nil {
		res.FailWithMsg(fmt.Sprintf("创建广告失败：%v", err), c)
		return
	}

	res.OkWithMsg("新增广告成功", c)
}
