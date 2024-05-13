package utils

import (
	"blog_server/global"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"mime/multipart"
	"net"
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
func CheckFileIsImage(suffix string) string {
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

// GetValidMsg 返回结构体中的msg参数 （使用gin校验器时 取出tag中定义的msg）
func GetValidMsg(err error, obj any) string {
	// 使用的时候，需要传obj的指针
	getObj := reflect.TypeOf(obj)
	global.Log.Infof("查看传入的obj实际类型：%v", getObj.Kind())
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
	if err != nil {
		return err.Error()
	}
	return "未知错误"
}

func GetUserRealIP(c *gin.Context) string {
	// 获取客户端IP
	// 先尝试从 X-Real-IP 头部获取 IP
	ip := c.Request.Header.Get("X-Real-IP")
	if ip == "" {
		// 如果 X-Real-IP 不存在，尝试从 X-Forwarded-For 头部获取
		ip = c.Request.Header.Get("X-Forwarded-For")
		// X-Forwarded-For 可能包含多个 IP，选择最远端的客户端 IP
		if ip != "" {
			ips := strings.Split(ip, ",")
			for i := len(ips) - 1; i >= 0; i-- {
				ip = strings.TrimSpace(ips[i])
				if net.ParseIP(ip) != nil {
					break
				}
			}
		}
	}

	// 如果所有头部都无效，使用 RemoteAddr
	if ip == "" {
		ip = c.Request.RemoteAddr
	}
	return ip

}
