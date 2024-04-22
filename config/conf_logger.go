package config

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Direct       string `yaml:"direct"`
	ShowLine     bool   `yaml:"showLine"`
	LogInConsole bool   `yaml:"logInConsole"`
}
