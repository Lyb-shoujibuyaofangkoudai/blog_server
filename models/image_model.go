package models

import (
	"blog_server/global"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"blog_server/utils"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"os"
)

type ImageModel struct {
	MODEL
	Path         string                 `json:"path"`                               // 图片路径
	Hash         string                 `json:"hash"`                               // 图片hash
	Name         string                 `gorm:"size:128" json:"name"`               // 图片名称
	Suffix       string                 `gorm:"size:8" json:"suffix"`               // 文件后缀
	Type         string                 `gorm:"size:8;default:'image'" json:"type"` // 文件类型 image 或者 file
	FileLocation ctype.FileLocationType `gorm:"size:8;default:1" json:"file_location"`
}

// BeforeDelete 删除前的钩子
func (file *ImageModel) BeforeDelete(tx *gorm.DB) (err error) {
	global.Log.Infof("删除图片", file.FileLocation, ctype.Local)
	if file.FileLocation == ctype.Local {
		//	 删除本地文件
		global.Log.Infof("删除本地文件路径：%v", file.Path)
		err := os.Remove(file.Path)
		if err != nil {
			global.Log.Errorf("删除本地文件失败：%v", err)
			return err
		}
	}
	return
}

func FileHashToDb(file *multipart.FileHeader, saveFilePath string, suffix string) res.FileUpload {
	fileHeader, err := file.Open()
	if err != nil {
		global.Log.Error(err.Error())
	}
	byteData, readErr := io.ReadAll(fileHeader)
	if readErr != nil {
		global.Log.Error(readErr.Error())
	}
	var imgModel ImageModel
	md5EncryptStr := utils.Md5(byteData)
	err = global.DB.Take(&imgModel, "hash = ?", md5EncryptStr).Error
	if err == nil {
		return res.FileUpload{
			FileName:  file.Filename,
			Url:       imgModel.Path,
			IsSuccess: false,
			ErrMsg:    "文件已存在",
		}
	}
	imgModel.Hash = md5EncryptStr
	imgModel.Name = file.Filename
	imgModel.Path = saveFilePath
	imgModel.Suffix = suffix
	imgModel.Type = utils.CheckFileIsImage(suffix)
	err = global.DB.Create(&imgModel).Error
	if err != nil {
		global.Log.Error(err.Error())
	}
	return res.FileUpload{
		FileName:  file.Filename,
		Url:       imgModel.Path,
		IsSuccess: true,
		ErrMsg:    "",
	}
}
