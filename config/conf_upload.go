package config

type Upload struct {
	Size   float64  `yaml:"size"`   // 单位MB 上传文件大小限定
	Path   string   `yaml:"path"`   // 上传文件路径
	Suffix []string `yaml:"suffix"` // 允许上传文件后缀
}
