/**
 * @date: 2022/7/26
 * @desc:
 */

package main

import (
	"github.com/daniuEvan/mygithub/core"
	"github.com/daniuEvan/mygithub/initialize"
	"github.com/gin-gonic/gin"
)

func main() {
	// 项目初始化
	var router *gin.Engine = initialize.InitStep()
	// 项目启动
	core.ListenAndServe(router)
}
