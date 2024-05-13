package flag

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/utils"
	"fmt"
)

func CreateUser(role string) {
	//	 创建用户

	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)
	fmt.Println("请输入用户名:")
	fmt.Scan(&userName)
	fmt.Println("请输入昵称:")
	fmt.Scan(&nickName)
	for {
		fmt.Println("请输入密码:")
		fmt.Scan(&password)
		fmt.Println("请再次输入密码:")
		fmt.Scan(&rePassword)
		if password == rePassword {
			break
		}
		fmt.Println("两次密码不一致")
	}
	fmt.Println("请输入邮箱:")
	fmt.Scan(&email)

	var userModel models.UserModel
	if global.DB.Take(&userModel, "email = ?", email).Error == nil {
		//	用户已存在
		fmt.Println("用户已存在")
		return
	}

	salt := utils.GenerateSalt(12) // 生成盐
	passwordMD5 := utils.EncryptPassword(password, salt)
	userModel = models.UserModel{
		UserName:   userName,
		Password:   passwordMD5,
		NickName:   nickName,
		Email:      email,
		Salt:       salt,
		SignStatus: ctype.SignEmail,
		Addr:       "内网注册地址",
		IP:         "127.0.0.1",
		Avatar:     "/static/avatar/avatar1.png", // 默认头像
	}
	if role == "admin" {
		userModel.Role = ctype.PermissionAdmin
	} else {
		userModel.Role = ctype.PermissionUser
	}
	if global.DB.Create(&userModel).Error != nil {
		fmt.Println("创建用户失败")
	}
	global.Log.Infof(fmt.Sprintf("创建用户成功,用户名:%s,密码:%s", userName, password))
}
