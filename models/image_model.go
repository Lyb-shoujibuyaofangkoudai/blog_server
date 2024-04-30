package models

type ImageModel struct {
	MODEL
	Path   string `json:"path"`                               // 图片路径
	Hash   string `json:"hash"`                               // 图片hash
	Name   string `gorm:"size:128" json:"name"`               // 图片名称
	Suffix string `gorm:"size:8" json:"suffix"`               // 文件后缀
	Type   string `gorm:"size:8;default:'image'" json:"type"` // 文件类型 image 或者 file
}
