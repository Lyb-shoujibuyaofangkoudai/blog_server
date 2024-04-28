package config

import (
	"fmt"
	"reflect"
)

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `canGet:"siteInfo" yaml:"site_info"`
	QQ       QQ       `canGet:"qq" yaml:"qq"`
	Email    Email    `canGet:"email" yaml:"email"`
	Jwt      Jwt      `canGet:"jwt" yaml:"jwt"`
	QiNiu    QiNiu    `canGet:"qiNiu" yaml:"qi_niu"`
}

// CanGetSettings 获取可以给后台显示和修改的配置信息
// CanGetSettings 提取所有标记为 canGet 的字段值
func (c *Config) CanGetSettings() map[string]string {
	result := make(map[string]string)
	// 获取结构体实例的反射值
	t := reflect.TypeOf(*c)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag
		canGetStr := tag.Get("canGet")

		if canGetStr != "" {
			// 获取字段的值
			result[canGetStr] = field.Name

		}
	}
	return result
}

func (c *Config) GetSettingByName(settingName string) reflect.Value {
	val := reflect.ValueOf(*c)
	return val.FieldByName(settingName)
}

func (c *Config) SetValue(config interface{}, fieldName string, value interface{}) error {
	configVal := reflect.ValueOf(config).Elem()
	field := configVal.FieldByName(fieldName)

	if !field.IsValid() {
		return fmt.Errorf("no such field: %s in Config", fieldName)
	}

	if !field.CanSet() {
		return fmt.Errorf("cannot set %s field of Config", fieldName)
	}

	field.Set(reflect.ValueOf(value))

	return nil
}
