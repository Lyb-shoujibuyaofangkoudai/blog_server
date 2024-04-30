package common

import (
	"blog_server/global"
	"blog_server/models"
)

type Option struct {
	models.Page
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	// 注意：这里的Find方法（采用了结构体查询方式）中传入的要是&list，如果传入的是&model，会将model进行修改，将里面的ID修改为获取的第一行数据的ID，会影响到后续的列表查询
	global.DB.Select("id").Where(&model).Find(&list).Count(&count)
	offset := option.Limit * (option.Page.Page - 1)
	if offset < 0 {
		offset = 0
	}
	var imageList []T
	err = global.DB.Select("name", "path").Where(&model).Limit(option.Limit).Offset(offset).Find(&imageList).Error
	return imageList, count, err
}
