/**
 * @date: 2022/7/28
 * @desc:
 */

package awesomeSpider

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"strings"
)

func awesomeGoSpider(enterUrl string, logger *zap.Logger) (dataResSlice []DataRes, err error) {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	c.OnHTML(`#readme > div.Box-body.px-5.pb-5 > article `, func(e *colly.HTMLElement) {
		dataRes := DataRes{
			DevLangCategory: "awesome-go",
		}
		toAppend := false // 过滤开头和结尾
		e.ForEach("[dir=auto]", func(i int, element *colly.HTMLElement) {
			// 过滤开头和结尾
			if !toAppend && element.Name == "h2" && element.Text != "Contents" && len(dataResSlice) == 0 {
				toAppend = true
			}
			if toAppend && element.Text == "Conferences" {
				toAppend = false
				// 最后加入一次
				if len(dataRes.ProjectInfos) > 0 {
					dataResSlice = append(dataResSlice, dataRes)
				}
			}

			if element.Name == "h2" || element.Name == "h3" {
				// 不为空 则放入结果切片
				if toAppend && len(dataRes.ProjectInfos) > 0 {
					dataResSlice = append(dataResSlice, dataRes)
				}
				dataRes = DataRes{
					DevLangCategory: "awesome-go",
					PurposeCategory: element.Text,
				}
			}
			switch element.Name {
			case "p": // 获取描述信息
				if !strings.HasSuffix(element.Text, " back to top") && toAppend {
					dataRes.PurposeCategoryDesc = element.Text
				}
			case "ul": // 获取 项目标签
				if toAppend {
					element.ForEach("li", func(j int, liElement *colly.HTMLElement) {
						addr := liElement.ChildAttr("a", "href")
						//projectName := liElement.ChildText("a")
						var projectName, projectDesc string
						if strings.Contains(liElement.Text, " - ") {
							projectName = strings.SplitN(liElement.Text, " - ", 2)[0]
							projectDesc = strings.SplitN(liElement.Text, " - ", 2)[1]
						} else {
							projectName = liElement.ChildText("a")
							projectDesc = liElement.Text
						}
						projectInfos := projectInfo{
							ProjectName: projectName,
							Addr:        addr,
							ProjectDesc: projectDesc,
						}
						dataRes.ProjectInfos = append(dataRes.ProjectInfos, projectInfos)
					})
				}
			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		logger.Info("OnRequest :", zap.String("url", r.URL.String()))
		//fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		logger.Info("OnResponse :", zap.String("url", fmt.Sprintf("Response %s: %d bytes", r.Request.URL, len(r.Body))))
	})

	c.OnError(func(r *colly.Response, err error) {
		logger.Error("OnError", zap.String("error", err.Error()))
	})

	err = c.Visit(enterUrl)
	return dataResSlice, err

}
