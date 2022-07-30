/**
 * @date: 2022/7/30
 * @desc:
 */

package spiderRouter

import (
	v1 "github.com/daniuEvan/mygithub/internal/controllor/v1"
	"github.com/gin-gonic/gin"
)

type AwesomeRouter struct{}

func (a *AwesomeRouter) InitAwesomeRouterV1(Router *gin.RouterGroup) {
	spiderRouter := Router.Group("spider")
	controller := v1.ControllerGroupApp.SpiderControllerGroup.AwesomeController
	{
		spiderRouter.GET("awesome/:category", controller.GetGithubData) // awesome 爬虫
	}
}
