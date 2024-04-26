package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primary_key" json:"id"` // 主键ID
	CreatedAt time.Time `json:"created_at"`            // 创建时间
	UpdatedAt time.Time `json:"updated_at"`            // 更新时间
}
