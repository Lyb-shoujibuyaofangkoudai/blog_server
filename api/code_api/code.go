package code_api

import (
	"blog_server/global"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type Email struct {
	Email string `form:"email" binding:"required,email" msg:"邮箱格式不正确"`
}

type Phone struct {
	Phone string `form:"phone" binding:"required,phone" msg:"手机号码格式不正确"`
}

func (CodeApi) Code(c *gin.Context) {
	param := c.Param("type")
	global.Log.Info("type参数为：" + param)
	if param == "email" {
		var email Email
		err := c.ShouldBindQuery(&email)
		if err != nil {
			res.FailWithValidateError(err, &email, c)
			return
		}
		GetEmailCode(c, email.Email)
	} else if param == "phone" {
	} else {
		res.FailWithMsg("type参数错误只能为email或phone", c)
	}
}
