/**
 * @date: 2022/7/28
 * @desc:
 */

package initialize

import (
	"github.com/daniuEvan/mygithub/global"
	"github.com/robfig/cron/v3"
)

func initSpiderCron() {
	c := cron.New()
	spiderConfig := global.ServerConfig.SpiderInfo
	// awesome spider
	entryId, err := c.AddJob(spiderConfig.AwesomeGo)
}
