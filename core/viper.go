package core

import (
	"flag"
	"fmt"
	"os"

	"gin-template/constants"
	"gin-template/global"
	_ "gin-template/packfile"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitViper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(constants.ConfigEnv); configEnv == "" {
				config = constants.ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", constants.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用MD_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func InitViper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}

	global.VIPER = v
	return v
}
