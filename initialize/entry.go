/**
 * @date: 2022/7/26
 * @desc:
 */

package initialize

import (
	"github.com/daniuEvan/mygithub/global"
	"github.com/gin-gonic/gin"
	"log"
)

func InitStep() (router *gin.Engine) {
	// 初始化 配置文件
	initConfigFromYaml()

	// 初始化 zap-logger
	initLogger()
	initSpiderLogger()

	// 初始化 翻译器
	err := initTrans(global.ServerConfig.Language)
	if err != nil {
		log.Fatalf("翻译器初始化失败:%s", err.Error())
	}

	// 初始化 自定义校验器
	err = initCustomValidator()
	if err != nil {
		log.Fatalf("自定义校验器初始化失败:%s", err.Error())
	}

	// 初始化 爬虫定时器
	initSpiderCron()

	// 初始化 路由
	router = initRouters()
	return router
}
