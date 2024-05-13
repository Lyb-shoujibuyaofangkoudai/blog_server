package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func phone(fl validator.FieldLevel) bool {
	pattern := `^(13[0-9]|14[5-9]|15[0-3,5-9]|16[6]|17[0-8]|18[0-9]|19[8,9])\d{8}$`
	phoneNum := fl.Field().String()
	// 编译正则表达式
	re := regexp.MustCompile(pattern)
	// 使用正则表达式的MatchString方法进行匹配
	return re.MatchString(phoneNum)
}

func code(fl validator.FieldLevel) bool {
	parent := fl.Parent()
	signInMethod := parent.FieldByName("SignInMethod").String()
	switch signInMethod {
	case "phone", "email":
		pattern := `^[0-9]{6}$`
		code := fl.Field().String()
		re := regexp.MustCompile(pattern)
		return re.MatchString(code)
	default:
		return true
	}
}
