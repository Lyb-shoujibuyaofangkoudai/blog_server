package user_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"blog_server/utils"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
	"time"
)

type UserInfo struct {
	UserName   string `form:"user_name" binding:"required,gte=3,lte=36" msg:"3<=用户名长度<=36"`
	NickName   string `form:"nick_name" binding:"gte=3,lte=36" msg:"3<=昵称长度<=36"`
	Password   string `form:"password" binding:"required"`
	RePassword string `form:"re_password" binding:"required,eqfield=Password" msg:"两次密码输入不同"`
	Avatar     string `form:"avatar"`
	Addr       string `form:"addr"`
	SignStatus int    `form:"sign_type" binding:"required,oneof=4 5" msg:"注册方式不能为空且仅支持4 5"`
}

type UserPhone struct {
	Phone string `form:"phone" binding:"required,phone" msg:"手机号码格式不正确"`
	Code  string `form:"code" binding:"required" msg:"验证码不能为空"`
}

type UserEmail struct {
	Email string `form:"email" binding:"required,email" msg:"邮箱格式不正确"`
	Code  string `form:"code" binding:"required" msg:"验证码不能为空"`
}

// 注册方式映射 采用策略模式
var signTypeMap = map[int]func(c *gin.Context) (any, string, error){
	int(ctype.SignQQ): func(c *gin.Context) (any, string, error) {
		global.Log.Info("QQ注册")
		return nil, "", nil
	},
	int(ctype.SignGitee): func(c *gin.Context) (any, string, error) {
		global.Log.Info("Gitee注册")
		return nil, "", nil
	},
	int(ctype.SignWechat): func(c *gin.Context) (any, string, error) {
		global.Log.Info("微信注册")
		return nil, "", nil
	},
	int(ctype.SignEmail): func(c *gin.Context) (any, string, error) {
		var userEmail UserEmail
		err := c.ShouldBind(&userEmail)
		if err != nil {
			res.FailWithValidateError(err, &userEmail, c)
			return nil, "", err
		}
		code := global.Redis.Get(context.Background(), userEmail.Email).Val()
		global.Log.Infof("redis中的验证码为：%s", code)
		if code == "" {
			res.FailWithMsg("验证码已过期，请重新获取验证码", c)
			return nil, "", errors.New("验证码已过期")
		} else if userEmail.Code != code {
			res.FailWithMsg("验证码错误", c)
			return nil, "", errors.New("验证码错误")
		}

		return userEmail, "userEmail", nil
	},
	int(ctype.SignPhone): func(c *gin.Context) (any, string, error) {
		var userPhone UserPhone
		err := c.ShouldBind(&userPhone)
		if err != nil {
			res.FailWithValidateError(err, &userPhone, c)
			return nil, "", err
		}
		return userPhone, "userPhone", nil
	},
}

// Register todo: 待完善注册功能 QQ GITEE 手机 微信 注册未完善
func (UserApi) Register(c *gin.Context) {

	var userInfo UserInfo
	err := c.ShouldBind(&userInfo)
	if err != nil {
		res.FailWithValidateError(err, &userInfo, c)
		return
	}
	cr, typeName, crErr := signTypeMap[userInfo.SignStatus](c)
	if crErr != nil {
		return
	}
	var userModel models.UserModel

	if cr != nil {
		if typeName == "userEmail" {
			userModel.Email = cr.(UserEmail).Email
			userModel.SignStatus = ctype.SignEmail
		} else if typeName == "userPhone" {
			userModel.Tel = cr.(UserPhone).Phone
			userModel.SignStatus = ctype.SignPhone
		}
	}

	affected := global.DB.Where(&userModel).First(&userModel).RowsAffected
	if affected > 0 {
		res.OkWithMsg("该用户已存在", c)
		return
	}
	// 使用加密安全的随机数生成器
	src := rand.NewSource(uint64(time.Now().UnixNano()))
	r := rand.New(src)
	// 生成1到62之间的随机整数
	seed := r.Intn(62) + 1
	// 生成盐值
	salt := utils.GenerateSalt(seed)

	userModel.Salt = salt
	// 密码加密
	userModel.Password = utils.EncryptPassword(userInfo.Password, salt)
	userModel.Role = ctype.PermissionUser
	userModel.IP = utils.GetUserRealIP(c)
	userModel.Avatar = "/static/avatar/avatar1.png"
	userModel.UserName = userInfo.UserName
	userModel.NickName = userInfo.NickName
	err = global.DB.Create(&userModel).Error
	if err != nil {
		global.Log.Errorf("注册失败: %v", err)
		res.FailWithMsg("注册失败", c)
		return
	}
	res.OkWithMsg(fmt.Sprintf("用户（%v）注册成功", userModel.UserName), c)
}
