package models

type MenuImageModel struct {
	MenuID     uint       `json:"menu_id"`
	ImageID    uint       `json:"image_id"`
	MenuModel  MenuModel  `gorm:"foreignKey:MenuID" json:"menu_model"`
	ImageModel ImageModel `gorm:"foreignKey:ImageID" json:"image_model"`
	Sort       int        `gorm:"size:10" json:"sort"`
}
