package bootstrap

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"panda/global"
)

func InitializeConfig() *viper.Viper {
	// 默认config路径
	config := "config.yaml"
	// 通过环境变量配置的路径
	if configEev := os.Getenv("Panda_CONFIG"); configEev != "" {
		config = configEev
	}
	// 初始化viper
	v := viper.New()
	// config路径
	v.SetConfigFile(config)
	// config类型
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}
	// 开始监听配置
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&global.App.Config); err != nil {
			log.Fatal("重载配置出错", err)
		}
	})
	// 将配置赋值给全局变量
	if err := v.Unmarshal(&global.App.Config); err != nil {
		log.Fatal("将配置赋值给全局变量出错", err)
	}
	global.LogCfg.Log = global.App.Config.Log
	return v
}
