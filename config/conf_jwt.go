package config

type Jwt struct {
	Secret  string `yaml:"secret"`  // 秘钥
	Expires int    `yaml:"expires"` // 过期时间 单位小时
	Issuer  string `yaml:"issuer"`  // 颁发人
}
