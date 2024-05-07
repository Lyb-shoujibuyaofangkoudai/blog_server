package images_api

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"github.com/gin-gonic/gin"
)

type FileUri struct {
	Type string `json:"type" uri:"type" binding:"oneof=image file all" msg:"params参数只能为image或file或all"` // oneof枚举校验器，type只能为image或者file
}

// FilesListView 获取文件列表
// @Tags 图片管理
// @Summary 获取文件列表
// @Description 获取文件列表
// @Accept application/json
// @Produce application/json
// @Param type path string true "文件类型"
// @Param page query int false "页码"
// @Param limit query int false "每页数量"
// @Param sort query string false "排序"
// @Success 200 {object}
func (ImagesApi) FilesListView(c *gin.Context) {
	fileTypeParams := FileUri{}
	err := c.ShouldBindUri(&fileTypeParams)
	if err != nil {
		res.FailWithValidateError(err, &fileTypeParams, c)
		return
	}
	pageQuery := models.Page{}
	if err := c.ShouldBindQuery(&pageQuery); err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}
	imageModel := models.ImageModel{}
	if fileTypeParams.Type != "all" {
		imageModel.Type = fileTypeParams.Type
	}
	list, count, err := common.ComList(imageModel, common.Option{
		Page: pageQuery,
		//Select: []string{"name", "path", "suffix", "type"},
	})
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OkWithList(list, count, c)
}
