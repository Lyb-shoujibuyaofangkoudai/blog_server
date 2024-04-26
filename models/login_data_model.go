package models

// LoginDataModel 用户登录信息表
type LoginDataModel struct {
	MODEL
	UserID    uint      `json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"_"`
	IP        string    `gorm:"size:20" json:"ip"` // 登录的IP
	NickName  string    `gorm:"size:42" json:"nick_name"`
	Token     string    `gorm:"size:128" json:"token"`
	Device    string    `gorm:"size:128" json:"device"`
	Addr      string    `gorm:"size:128" json:"addr"`
}
