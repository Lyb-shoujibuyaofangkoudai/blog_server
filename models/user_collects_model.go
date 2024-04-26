package models

import "time"

// UserCollectsModel 记录用户什么时候收藏了什么文章
type UserCollectsModel struct {
	UserID    uint      `gorm:"primaryKey"`
	ArticleID uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	// 下面不会在表中生成对应字段的
	UserModel    UserModel    `gorm:"foreignKey:UserID"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
}
