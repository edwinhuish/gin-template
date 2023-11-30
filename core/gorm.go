package core

import (
	"gin-template/global"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func InitGorm() {
	global.DB = InitGormMysql()
}
