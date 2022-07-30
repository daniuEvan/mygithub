/**
 * @date: 2022/7/30
 * @desc: 爬虫service
 */

package spiderService

import (
	"fmt"
	"github.com/daniuEvan/mygithub/global"
	"github.com/daniuEvan/mygithub/spider/awesomeSpider"
)

type AwesomeService struct{}

func (a *AwesomeService) AwesomeSpider(langCategory string) error {
	dataResSlice, err := awesomeSpider.AwesomeSpider(langCategory, global.Logger)
	fmt.Println(dataResSlice)
	return err
}
