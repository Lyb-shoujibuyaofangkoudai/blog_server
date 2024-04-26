package config

type SiteInfo struct {
	CreatedAt     string `yaml:"created_at" json:"created_at"`
	BeiAn         string `yaml:"bei_an" json:"bei_an"`
	Title         string `yaml:"title" json:"title"`
	WechatQrImage string `yaml:"wechat_qr_image" json:"wechat_qr_image"`
	QqQrImage     string `yaml:"qq_qr_image" json:"qq_qr_image"`
	QqNumber      string `yaml:"qq_number" json:"qq_number"`
	Version       string `yaml:"version" json:"version"`
	Email         string `yaml:"email" json:"email"`
	Github        string `yaml:"github" json:"github"`
	Gitee         string `yaml:"gitee" json:"gitee"`
	Name          string `yaml:"name" json:"name"`
	Job           string `yaml:"job" json:"job"`
	Address       string `yaml:"address" json:"address"`
	Slogan        string `yaml:"slogan" json:"slogan"`
	SloganEn      string `yaml:"slogan_en" json:"slogan_en"`
}
