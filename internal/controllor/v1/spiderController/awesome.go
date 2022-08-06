/**
 * @date: 2022/7/30
 * @desc:
 */

package spiderController

import (
	"github.com/daniuEvan/mygithub/common/response"
	"github.com/daniuEvan/mygithub/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type AwesomeController struct {
}

func (a *AwesomeController) GetGithubData(ctx *gin.Context) {
	langCategory := ctx.Param("category")
	switch langCategory {
	case "python":
		err := awesomeService.AwesomeSpider("awesome-" + langCategory)
		if err != nil {
			global.Logger.Error("awesome-python service error.", zap.String("error msg", err.Error()))
			response.Response(ctx, http.StatusInternalServerError, 500, nil, "awesome-python service error.")
			return
		}
		response.Success(ctx, nil, "awesome-python execute success")
	case "go":
		err := awesomeService.AwesomeSpider("awesome-" + langCategory)
		if err != nil {
			global.Logger.Error("awesome-go service error.", zap.String("error msg", err.Error()))
			response.Response(ctx, http.StatusInternalServerError, 500, nil, "awesome-go service error.")
			return
		}
		response.Success(ctx, nil, "awesome-go execute success")
	default:
		global.Logger.Error("awesome spider category error!", zap.String("langCategory", langCategory))
		response.Fail(ctx, nil, "awesome spider category error!")
	}
}
