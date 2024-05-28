package core

import (
	"blog_server/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", global.Config.Redis.Host, global.Config.Redis.Port),
		Password:     global.Config.Redis.Password, // 没有密码，默认值
		DB:           global.Config.Redis.DB,       // 默认DB 0
		PoolSize:     global.Config.Redis.PoolSize,
		MinIdleConns: global.Config.Redis.MiniDleConns,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		global.Log.Errorf("初始化redis失败：%v", err.Error())
		return nil
	}
	global.Log.Info("初始化redis成功")
	return rdb
}
