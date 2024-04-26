package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	QQ       QQ       `yaml:"qq"`
	Email    Email    `yaml:"email"`
	Jwt      Jwt      `yaml:"jwt"`
	QiNiu    QiNiu    `yaml:"qi_niu"`
}
