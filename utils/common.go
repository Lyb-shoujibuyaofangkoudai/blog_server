package utils

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"mime/multipart"
	"path"
	"reflect"
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

// CheckFileIsImage 校验文件是否是图片
func checkFileIsImage(suffix string) string {
	imageSuffix := []string{"jpg", "jpeg", "png", "gif", "bmp", "webp", "svg", "tiff"}
	for _, image := range imageSuffix {
		if image == suffix {
			return "image"
		}
	}
	return "file"
}

// CheckFileSuffixIsRight 校验文件后缀是否符合配置
func CheckFileSuffixIsRight(file *multipart.FileHeader) (suffix string, err error) {
	fileSuffix := strings.ToLower(strings.Split(file.Filename, ".")[len(strings.Split(file.Filename, "."))-1])
	for _, suffix := range global.Config.Upload.Suffix {
		if suffix == fileSuffix {
			return fileSuffix, nil
		}
	}
	return "", errors.New(fmt.Sprintf("不允许上传%s的文件", fileSuffix))
}

func GenerationFilePath(fileName string) string {
	return path.Join(global.Config.Upload.Path, fmt.Sprintf("%d", time.Now().UnixNano())+"_"+fileName)
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
	imgModel.Suffix = suffix
	imgModel.Type = checkFileIsImage(suffix)
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

// GetValidMsg 返回结构体中的msg参数 （使用gin校验器时 取出tag中定义的msg）
func GetValidMsg(err error, obj any) string {
	// 使用的时候，需要传obj的指针
	getObj := reflect.TypeOf(obj)
	// 将err接口断言为具体类型
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误信息
			// 根据报错字段名，获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}

	return err.Error()
}
