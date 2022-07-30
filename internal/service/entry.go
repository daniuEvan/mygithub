/**
 * @date: 2022/7/30
 * @desc:
 */

package service

import (
	"github.com/daniuEvan/mygithub/internal/service/spiderService"
)

type serviceGroup struct {
	// 爬虫 service
	SpiderServiceGroup spiderService.ServiceGroup
}

var ServiceGroupApp = new(serviceGroup)
