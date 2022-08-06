/**
 * @date: 2022/7/28
 * @desc:
 */

package awesomeSpider

import (
	"errors"
	"github.com/daniuEvan/mygithub/internal/model/spiderModel"
	"github.com/daniuEvan/mygithub/spider/getStarCount"
	"go.uber.org/zap"
	"regexp"
	"strings"
	"time"
)

func AwesomeSpider(langCategory string, logger *zap.Logger) (dataResSlice []DataRes, err error) {
	entryAddr, _ := awesomeDatasource[langCategory]
	switch langCategory {
	case "awesome-python":
		dataResSlice, err = awesomePythonSpider(entryAddr, logger)
	case "awesome-go":
		dataResSlice, err = awesomeGoSpider(entryAddr, logger)
	default:
		return nil, errors.New("爬虫类型错误")
	}
	// todo 过滤没有项目地址不是github的
	for _, dataRes := range dataResSlice {
		var (
			devLangCategory     = dataRes.DevLangCategory
			purposeCategory     = dataRes.PurposeCategory
			purposeCategoryDesc = dataRes.PurposeCategoryDesc
			projectInfos        = dataRes.ProjectInfos
		)
		for _, pInfo := range projectInfos {
			time.Sleep(500 * time.Millisecond)

			var (
				projectName = pInfo.ProjectName
				addr        = pInfo.Addr
				projectDesc = pInfo.ProjectDesc
			)

			var awesome = &spiderModel.Awesome{
				DevLangCategory:     devLangCategory,
				PurposeCategory:     purposeCategory,
				PurposeCategoryDesc: purposeCategoryDesc,
				ProjectName:         projectName,
				Addr:                addr,
				ProjectDesc:         projectDesc,
			}
			if !strings.Contains(addr, "://github.com/") {
				// 不是github托管 直接入库
				println("projectName:", projectName, "addr:", addr)
				err = toDatabase(awesome)
				if err != nil {
					return nil, err
				}
				continue
			}
			awesome.IsGithub = 1
			regexpRule := "https?://github.com/(.*?)/"
			r := regexp.MustCompile(regexpRule)
			awesome.Author = r.FindStringSubmatch(addr)[1]
			// 获取星星,作者,最后提交时间
			starNum, commitTime, err := getStarCount.GetStarCount(addr, logger)
			if err != nil {
				return nil, err
			}
			awesome.StarCount = starNum
			awesome.LastCommitTime = commitTime
			// 入库
			err = toDatabase(awesome)
			if err != nil {
				return nil, err
			}
		}
	}

	return dataResSlice, err
}
