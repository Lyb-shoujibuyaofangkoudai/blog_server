package images_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// FileUploadView 上传单个文件
func (ImagesApi) FileUploadView(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		global.Log.Errorf("FormFile: %v", err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	suffix, suffixErr := utils.CheckFileSuffixIsRight(file)
	if suffixErr != nil {
		global.Log.Errorf("CheckFileSuffixIsRight: %v", suffixErr)
		res.FailWithMsg(suffixErr.Error(), c)
		return
	}
	isLimitNotExceeded := utils.CheckFileSizeIsRight(float64(file.Size))
	if !isLimitNotExceeded {
		global.Log.Errorf("CheckFileSizeIsRight: %v", err)
		res.FailWithCode(res.FileSizeExceeded, c)
		return
	}

	filePath := utils.GenerationFilePath(file.Filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	imgRes := models.FileHashToDb(file, filePath, suffix)
	res.OkWithData(imgRes, c)
}

// FilesUploadViews 上传多个文件
func (ImagesApi) FilesUploadViews(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	var resFileList []res.FileUpload
	files := form.File["files"]
	global.Log.Infof("本次上传文件数量为%v", len(files))
	for _, file := range files {
		suffix, suffixErr := utils.CheckFileSuffixIsRight(file)
		if suffixErr != nil {
			resFileList = append(resFileList, res.FileUpload{
				FileName:  file.Filename,
				Url:       "",
				IsSuccess: false,
				ErrMsg:    suffixErr.Error(),
			})
			continue
		}
		if !utils.CheckFileSizeIsRight(float64(file.Size)) {
			// 超出文件大小限制
			resFileList = append(resFileList, res.FileUpload{
				FileName:  file.Filename,
				Url:       "",
				IsSuccess: false,
				ErrMsg:    fmt.Sprintf("当前文件大小为%vM，已超出%vM限制", strconv.FormatFloat(float64(file.Size)/float64(1024*1024), 'f', 2, 64), strconv.FormatFloat(global.Config.Upload.Size, 'f', 2, 64)),
			})
		} else {
			filePath := utils.GenerationFilePath(file.Filename)
			err = c.SaveUploadedFile(file, filePath)
			if err != nil {
				resFileList = append(resFileList, res.FileUpload{
					FileName:  file.Filename,
					Url:       "",
					IsSuccess: false,
					ErrMsg:    err.Error(),
				})
			} else {
				imgRes := models.FileHashToDb(file, filePath, suffix)
				resFileList = append(resFileList, imgRes)
			}
		}
	}
	res.OkWithData(resFileList, c)
}
