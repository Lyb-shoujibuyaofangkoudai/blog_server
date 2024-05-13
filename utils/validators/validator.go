package validators

import (
	"blog_server/global"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// GetValidMsg
// 获取参数校验触发校验器失败，获取tag中msg字段的信息
func GetValidMsg(err error, obj any) string {
	global.Log.Errorf("错误：%v", err.Error())
	// 使用的时候，需要传obj的指针
	getObj := reflect.TypeOf(obj)
	// 将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误信息
			// 根据报错字段名，获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				if msg != "" {
					return msg
				}
			}
		}
	}
	return err.Error()
}
