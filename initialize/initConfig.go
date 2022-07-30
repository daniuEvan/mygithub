/*
 * @date: 2021/12/15
 * @desc: ...
 */

package initialize

import (
	"fmt"
	"github.com/daniuEvan/mygithub/global"
	"github.com/daniuEvan/mygithub/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func initConfigFromYaml() {
	debugEnv := utils.IsDebugEnv()
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("./%s-dev.yaml", configFilePrefix)
	if !debugEnv {
		configFileName = fmt.Sprintf("./%s.yaml", configFilePrefix)
	}
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("配置初始化失败:%s", err.Error())
	}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		log.Fatalf("配置初始化失败:%s", err.Error())
	}
	log.Println("配置文件初始化完成.")
	if debugEnv {
		log.Println(global.ServerConfig)
	}
	// 监测配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		log.Println("配置文件重新初始化:")
		if err := v.ReadInConfig(); err != nil {
			log.Fatalf("配置重新初始化失败:%s", err.Error())
		}
		if err := v.Unmarshal(global.ServerConfig); err != nil {
			log.Fatalf("配置重新初始化失败:%s", err.Error())
		}
		log.Println("配置文件重新初始化完成.")
	})

}
