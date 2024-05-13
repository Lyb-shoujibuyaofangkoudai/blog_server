package models

import "blog_server/models/ctype"

// UserModel 用户表
type UserModel struct {
	MODEL
	NickName      string           `gorm:"size:42" json:"nick_name"`                                                                            // 昵称
	UserName      string           `gorm:"size:36" json:"user_name"`                                                                            // 用户名
	Password      string           `gorm:"size:128" json:"password"`                                                                            // 密码
	Salt          string           `json:"salt"`                                                                                                // 密码盐
	Avatar        string           `gorm:"size:26" json:"avatar"`                                                                               // 头像
	Email         string           `gorm:"size:128" json:"email"`                                                                               // 邮箱
	Tel           string           `gorm:"size:18" json:"tel"`                                                                                  // 手机号
	Addr          string           `gorm:"size:64" json:"address"`                                                                              // 地址
	Token         string           `gorm:"size:64" json:"token"`                                                                                // token
	IP            string           `gorm:"size:64" json:"ip"`                                                                                   // ip
	Role          ctype.Role       `gorm:"size:4;default:1" json:"role"`                                                                        // 角色 1 管理员 2 普通用户 3 游客 4 被禁用
	SignStatus    ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"`                                                                 // 注册来源 qq 邮箱 手机等
	ArticleModel  []ArticleModel   `gorm:"foreignKey:UserID" json:"ArticleModels"`                                                              // 发布文章列表
	CollectsModel []ArticleModel   `gorm:"many2many:user_collects_models;joinForeignKey:UserID;joinReferences:ArticleID" json:"CollectsModels"` // 收藏文章列表
}

//func ParseRole(role Role) string {
//	switch role {
//	case PermissionAdmin:
//		return "管理员"
//
//	case PermissionUser:
//		return "普通用户"
//
//	case PermissionVisitor:
//		return "游客"
//
//	case PermissionSuperDisabled:
//		return "被禁用"
//
//	default:
//		return "其他"
//	}
//}
