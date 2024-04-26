package models

import "blog_server/models/ctype"

type ArticleModel struct {
	MODEL
	Title        string         `gorm:"size:64" json:"title"`                       // 文章标题
	Abstract     string         `json:"abstract"`                                   // 文章摘要
	Content      string         `gorm:"type:longtext" json:"content"`               // 文章内容
	LookCount    int            `gorm:"default:0" json:"look_count"`                // 浏览量
	CommentCount int            `gorm:"default:0" json:"comment_count"`             // 评论量
	DiggCount    int            `gorm:"default:0" json:"digg_count"`                // 点赞量
	TagModel     []TagModel     `gorm:"many2many:article_tag" json:"tag_model"`     // 文章标签
	CommentModel []CommentModel `gorm:"foreignKey:ArticleID" json:"comment_models"` // 文章评论列表
	UserModel    UserModel      `gorm:"foreignKey:UserID" json:"user_model"`        // 文章作者
	UserID       uint           `json:"user_id"`                                    // 用户
	ArticleID    uint           `json:"article_id"`                                 // 文章ID
	Category     string         `gorm:"size:20" json:"category"`                    // 文章分类
	Source       string         `json:"source"`                                     // 文章来源
	Link         string         `json:"link"`                                       // 文章来源链接
	Cover        ImageModel     `json:"cover"`                                      // 文章封面
	CoverID      uint           `json:"cover_id"`                                   // 文章封面ID
	NickName     string         `gorm:"size:42" json:"nick_name"`                   // 文章作者昵称
	CoverPath    string         `json:"cover_path"`                                 // 文章封面路径
	Tags         ctype.Array    `gorm:"string;size:64" json:"tags"`                 // 文章标签
}
