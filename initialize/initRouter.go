/*
 * @date: 2021/12/15
 * @desc: ...
 */

package initialize

import (
	"github.com/daniuEvan/mygithub/global"
	"github.com/daniuEvan/mygithub/internal/router"
	"github.com/daniuEvan/mygithub/middleware"
	"github.com/daniuEvan/mygithub/utils"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// initRouters 初始化gin 路由
func initRouters() *gin.Engine {
	var defaultEngine *gin.Engine
	if utils.IsDebugEnv() {
		defaultEngine = gin.Default()
		pprof.Register(defaultEngine)
	} else {
		gin.SetMode(gin.ReleaseMode)
		defaultEngine = gin.New()
	}
	defaultEngine.Use(middleware.Cors(), middleware.GinLogger(global.Logger)) // 跨域

	spiderRouter := router.RouterGroupApp.SpiderRouterGroup

	// 方便统一路由
	apiGroup := defaultEngine.Group("api")
	v1 := apiGroup.Group("v1")
	{
		spiderRouter.InitAwesomeRouterV1(v1)
	}
	return defaultEngine
}
