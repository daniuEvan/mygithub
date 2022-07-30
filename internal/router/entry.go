/**
 * @date: 2022/7/30
 * @desc:
 */

package router

import (
	"github.com/daniuEvan/mygithub/internal/router/spiderRouter"
)

type routerGroup struct {
	SpiderRouterGroup spiderRouter.RouterGroup
}

var RouterGroupApp = new(routerGroup)
