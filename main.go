package main

import (
	"time"

	"gin-template/core"
	"gin-template/global"

	_ "net/http/pprof"
)

func main() {

	time.Local = time.FixedZone("CST", 8*3600) // 东八区

	core.InitViper() // 初始化Viper
	core.InitZap()   // 初始化zap日志库

	if global.CONFIG.System.UseRedis {
		// 初始化redis服务
		core.InitRedis()
	}
	// core.InitGorm() // gorm连接数据库
	core.InitHttp()
}
