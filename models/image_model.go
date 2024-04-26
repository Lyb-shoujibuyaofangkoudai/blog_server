package models

type ImageModel struct {
	MODEL
	Path string `json:"path"`                // 图片路径
	Hash string `json:"hash"`                // 图片hash
	Name string `gorm:"size:32" json:"name"` // 图片名称

}
