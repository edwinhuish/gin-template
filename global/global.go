package global

import (
	"go.uber.org/zap"

	"gin-template/config"
	"gin-template/utils/timer"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	DBList map[string]*gorm.DB
	REDIS  *redis.Client
	CONFIG config.Server
	VIPER  *viper.Viper
	LOG    *zap.Logger
	TIMER  timer.Timer = timer.NewTimerTask()
)
