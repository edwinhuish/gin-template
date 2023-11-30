package core

import (
	"gin-template/global"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func InitRedis() {
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.LOG.Error("redis connect ping failed, err:", zap.Error(err))
		return
	}

	global.LOG.Info("redis connect ping response:", zap.String("pong", pong))
	global.REDIS = client
}
