package user_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type LoginInfo struct {
	LoginStatus string `form:"login_type" binding:"required,oneof=password phone_code" msg:"login_type为必填参数，登录方式只能为账号密码（password）或手机验证码登录（phone_code）"`
}

type PhoneCode struct {
	phone string `form:"phone" binding:"required,phone" msg:"手机号码格式不正确"`
	code  string `form:"code" binding:"required" msg:"验证码不能为空"`
}

type AccountPassword struct {
	Account  string `form:"account" binding:"required" msg:"账号不能为空，账号为手机或邮箱"`
	Password string `form:"password" binding:"required" msg:"密码不能为空"`
}

// 登录方式映射 使用策略模式
var loginTypeMap = map[string]func(c *gin.Context) (any, error){
	"password": func(c *gin.Context) (any, error) {
		global.Log.Info("账号密码登录")
		var accountPassword AccountPassword
		err := c.ShouldBind(&accountPassword)
		if err != nil {
			res.FailWithValidateError(err, &accountPassword, c)
			return "", errors.New("参数错误")
		}
		user := models.UserModel{}
		affected := global.DB.
			Where("tel = ?", accountPassword.Account).
			Or("email = ?", accountPassword.Account).
			Find(&user).RowsAffected
		global.Log.Infof("user: %v", affected)
		if affected == 0 {
			res.FailWithMsg("用户不存在", c)
			return nil, errors.New("用户不存在")
		}
		return user, nil
	},
	"phone_code": func(c *gin.Context) (any, error) {
		global.Log.Info("手机验证码登录")
		return nil, nil
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
	userInfo, err := loginTypeMap[loginInfo.LoginStatus](c)
	if err != nil {
		global.Log.Error(err)
		return
	}

	userKeyInRedis := fmt.Sprintf("%i_%s", userInfo.(models.UserModel).ID, userInfo.(models.UserModel).Salt) // 用户id+salt

	// 查看当前用户是否登录了
	currentUserToken := global.Redis.Get(context.Background(), userKeyInRedis).Val()

	if currentUserToken != "" {
		// 解析token
		tokenParse, _ := utils.ParseTokenRs256(currentUserToken)
		if tokenParse.ExpiresAt.Unix()-time.Now().Unix() > 60*60*2 {
			// 如果token时效大于4小时，则直接返回数据
			res.Ok(map[string]any{
				"token":    currentUserToken,
				"userInfo": userInfo,
			}, "用户已经登录过了，请勿重复登录", c)
			return
		}
	}

	// 生成token
	token, err := utils.GenerateTokenUsingRS256(userInfo.(models.UserModel).ID, userInfo.(models.UserModel).UserName)
	if err != nil {

		res.FailWithMsg("登录失败，请稍后再试", c)
		global.Log.Errorf("生成token失败: %v", err)
		return
	}

	global.Redis.Set(context.Background(), userKeyInRedis, token, time.Hour*time.Duration(global.Config.Jwt.Expires))

	res.OkWithData(map[string]any{
		"token":    token,
		"userInfo": userInfo,
	}, c)
}
