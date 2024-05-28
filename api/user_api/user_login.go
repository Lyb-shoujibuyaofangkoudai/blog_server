package user_api

import (
	"blog_server/global"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type LoginInfo struct {
	LoginStatus string `form:"login_type" binding:"required,oneof=password phone_code" msg:"登录方式只能为账号密码或手机验证码登录"`
}

type PhoneCode struct {
	phone string `form:"phone" binding:"required,phone" msg:"手机号码格式不正确"`
	code  string `form:"code" binding:"required" msg:"验证码不能为空"`
}

type AccountPassword struct {
	Account  string `form:"account" binding:"required" msg:"账号不能为空，账号为手机，邮箱或用户名"`
	Password string `form:"password" binding:"required" msg:"密码不能为空"`
}

// 登录方式映射
var loginTypeMap = map[string]func(c *gin.Context){
	"password": func(c *gin.Context) {
		global.Log.Info("账号密码登录")
	},
	"phone_code": func(c *gin.Context) {
		global.Log.Info("手机验证码登录")
	},
}

// Login 登录
func (UserApi) Login(c *gin.Context) {
	var loginInfo LoginInfo
	err := c.ShouldBind(&loginInfo)
	if err != nil {
		res.FailWithValidateError(err, &loginInfo, c)
		return
	}
	loginTypeMap[loginInfo.LoginStatus](c)
	res.OkWith(c)
}
