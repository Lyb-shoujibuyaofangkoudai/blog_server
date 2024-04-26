package models

type MenuImageModel struct {
	MenuID     uint       `json:"menu_id"`
	MenuModel  MenuModel  `gorm:"foreignKey:MenuID" json:"menu_model"`
	ImageID    uint       `json:"image_id"`
	ImageModel ImageModel `gorm:"foreignKey:ImageID" json:"image_model"`
	Sort       int        `gorm:"size:10" json:"sort"`
}
