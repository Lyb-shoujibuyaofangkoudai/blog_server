package config

type Email struct {
	Host              string `yaml:"host" json:"host"`
	Port              int    `yaml:"port" json:"port"`
	User              string `yaml:"user" json:"user"` // 发件人邮箱
	Password          string `yaml:"password" json:"password"`
	DefaultFromEmail  string `yaml:"default_from_email" json:"default_from_email"` // 默认发件人名字
	UserSSL           bool   `yaml:"user_ssl" json:"user_ssl"`                     // 是否使用ssl
	UserTls           bool   `yaml:"user_tls" json:"user_tls"`                     // 是否使用tls
	AuthorizationCode string `yaml:"authorization_code" json:"authorization_code"` // 授权码
}
