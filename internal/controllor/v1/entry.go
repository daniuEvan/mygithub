/**
 * @date: 2022/7/30
 * @desc:
 */

package v1

import "github.com/daniuEvan/mygithub/internal/controllor/v1/spiderController"

type controllerGroup struct {
	SpiderControllerGroup spiderController.ControllerGroup
}

var ControllerGroupApp = new(controllerGroup)
