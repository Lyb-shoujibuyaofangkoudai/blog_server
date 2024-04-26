package models

type TagModel struct {
	MODEL
	Title    string         `gorm:"size:32" json:"title"`           // 标签名
	Articles []ArticleModel `gorm:"many2many:article_tag" json:"_"` // 标签对应的文章
}
