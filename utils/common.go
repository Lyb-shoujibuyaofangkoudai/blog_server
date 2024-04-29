package utils

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"strings"
	"time"
)

// CheckFileSizeIsRight 检验上传文件大小是否符合配置
func CheckFileSizeIsRight(fileSize float64) bool {
	fileSizeToMb := fileSize / float64(1024*1024)
	if fileSizeToMb > global.Config.Upload.Size {
		return false
	}
	return true
}

// CheckFileSuffixIsRight 校验文件后缀是否符合配置
func CheckFileSuffixIsRight(file *multipart.FileHeader) ([]string, error) {
	fileSuffix := strings.Split(file.Filename, ".")[len(strings.Split(file.Filename, "."))-1]
	for _, suffix := range global.Config.Upload.Suffix {
		if suffix == strings.ToLower(fileSuffix) {
			return global.Config.Upload.Suffix, nil
		}
	}
	return global.Config.Upload.Suffix, errors.New(fmt.Sprintf("不允许上传%s的文件", fileSuffix))
}

func GenerationFilePath(fileName string) string {
	return path.Join(global.Config.Upload.Path, fmt.Sprintf("%d", time.Now().UnixNano())+"_"+fileName)
}

func FileHashToDb(file *multipart.FileHeader, saveFilePath string) res.FileUpload {
	fileHeader, err := file.Open()
	if err != nil {
		global.Log.Error(err.Error())
	}
	byteData, readErr := io.ReadAll(fileHeader)
	if readErr != nil {
		global.Log.Error(readErr.Error())
	}
	var imgModel models.ImageModel
	md5EncryptStr := Md5(byteData)
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
