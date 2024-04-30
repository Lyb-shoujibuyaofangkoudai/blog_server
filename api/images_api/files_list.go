package images_api

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"blog_server/utils"
	"github.com/gin-gonic/gin"
)

type FileUri struct {
	Type string `json:"type" uri:"type" binding:"oneof=image file" msg:"params参数只能为image或file"` // oneof枚举校验器，type只能为image或者file
}

func (ImagesApi) FilesListView(c *gin.Context) {
	fileTypeParams := FileUri{}
	if err := c.ShouldBindUri(&fileTypeParams); err != nil {
		msg := utils.GetValidMsg(err, &fileTypeParams)
		res.FailWithMsg(msg, c)
		return
	}
	pageQuery := models.Page{}
	if err := c.ShouldBindQuery(&pageQuery); err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}
	imageModel := models.ImageModel{
		Type: fileTypeParams.Type,
	}
	list, count, err := common.ComList(imageModel, common.Option{
		Page: pageQuery,
	})
	if err != nil {
		res.FailWithMsg("获取失败", c)
		return
	}
	res.OkWithList(list, count, c)
}
