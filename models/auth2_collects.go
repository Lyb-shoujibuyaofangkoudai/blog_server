package models

import "time"

// Auth2Collects 记录用户什么时候收藏了什么文章
type Auth2Collects struct {
	UserID       uint         `gorm:"primaryKey"`
	AuthModel    ArticleModel `gorm:"foreignKey:UserID"`
	ArticleID    uint         `gorm:"primaryKey"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
	CreatedAt    time.Time
}
