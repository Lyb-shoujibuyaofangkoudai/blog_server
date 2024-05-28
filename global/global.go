package global

import (
	"blog_server/config"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config   *config.Config
	DB       *gorm.DB
	Redis    *redis.Client
	Log      *logrus.Logger
	MysqlLog *logrus.Logger
)
