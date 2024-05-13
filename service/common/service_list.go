package common

import (
	"blog_server/global"
	"blog_server/models"
)

type RemoveFileList struct {
	Ids []uint `binding:"required" msg:"ids参数不能为空"`
}

type Option struct {
	models.Page
	Select []string
}

// ComSingleList 公用获取列表数据方法（适用于单表）
func ComSingleList[T any](model *T, option Option) (list []T, count int64, err error) {
	// 注意：这里的Find方法（采用了结构体查询方式）中传入的要是&list，如果传入的是&model，会将model进行修改，将里面的ID修改为获取的第一行数据的ID，会影响到后续的列表查询
	global.DB.Select("id").Where(&model).Find(&list).Count(&count)
	if option.Page.Page == 0 {
		option.Page.Page = 1
	}
	if option.Limit == 0 {
		option.Limit = 10
	}
	offset := option.Limit * (option.Page.Page - 1)
	if offset < 0 {
		offset = 0
	}
	var imageList []T
	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认按照创建时间倒序排序
	}
	err = global.DB.Select(option.Select).Order(option.Sort).Where(&model).Limit(option.Limit).Offset(offset).Find(&imageList).Error
	return imageList, count, err
}
