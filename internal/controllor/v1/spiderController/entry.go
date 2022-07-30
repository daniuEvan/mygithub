/**
 * @date: 2022/7/30
 * @desc:
 */

package spiderController

import "github.com/daniuEvan/mygithub/internal/service"

type ControllerGroup struct {
	AwesomeController
}

var (
	awesomeService = service.ServiceGroupApp.SpiderServiceGroup.AwesomeService
)
