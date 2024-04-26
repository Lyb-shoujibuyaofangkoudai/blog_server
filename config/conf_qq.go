package config

import "fmt"

type QQ struct {
	AppID    string `yaml:"app_id" json:"app_id"`
	Key      string `yaml:"key" json:"key"`
	Redirect string `yaml:"redirect" json:"redirect"`
}

func (q *QQ) GetQQImage() string {
	if q.Key == "" || q.AppID == "" || q.Redirect == "" {
		return ""
	}
	return fmt.Sprintf("")
}
