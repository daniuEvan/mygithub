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

func awesomePythonSpider(enterUrl string, logger *zap.Logger) (dataResSlice []DataRes, err error) {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	projectNames := ""
	c.OnHTML(`#readme > div.Box-body.px-5.pb-5 > article`, func(e *colly.HTMLElement) {
		dataRes := DataRes{
			DevLangCategory: "awesome-python",
		}
		toAppend := false // 过滤开头和结尾
		e.ForEachWithBreak("[dir=auto]", func(i int, element *colly.HTMLElement) bool {
			// 过滤开头和结尾
			if !toAppend && element.Name == "h2" && len(dataResSlice) == 0 {
				toAppend = true
			}
			if toAppend && element.Text == "Resources" {
				if len(dataRes.ProjectInfos) > 0 {
					dataResSlice = append(dataResSlice, dataRes)
				}
				return false
			}

			if element.Name == "h2" || element.Name == "h3" {
				// 不为空 则放入结果切片
				if toAppend && len(dataRes.ProjectInfos) > 0 {
					dataResSlice = append(dataResSlice, dataRes)
				}
				dataRes = DataRes{
					DevLangCategory: "awesome-python",
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
						ulDir := liElement.ChildAttr("ul", "dir")
						//fmt.Println(ulDir)
						if ulDir == "auto" {

							// 解析子 ul
							liElement.ForEach("ul[dir=auto] > li", func(l int, childLiElement *colly.HTMLElement) {

								projectInfo := toParseLi(childLiElement)
								projectName := projectInfo.ProjectName
								if !strings.Contains(projectName, "awesome") && !strings.Contains(projectNames, projectName) {
									dataRes.ProjectInfos = append(dataRes.ProjectInfos, projectInfo)
									projectNames += fmt.Sprintf("%s,", projectName)
								}

							})
						} else {

							projectInfo := toParseLi(liElement)
							projectName := projectInfo.ProjectName

							if !strings.Contains(projectName, "awesome") && !strings.Contains(projectNames, projectName) {
								dataRes.ProjectInfos = append(dataRes.ProjectInfos, projectInfo)
								projectNames += fmt.Sprintf("%s,", projectName)

							}
						}

					})
				}
			}
			return true
		})
	})

	c.OnRequest(func(r *colly.Request) {
		logger.Info("PythonSpiderOnRequest :", zap.String("url", r.URL.String()))
		//fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		//logger.Info("PythonSpiderOnResponse :", zap.String("url", fmt.Sprintf("Response %s: %d bytes", r.Request.URL, len(r.Body))))
	})

	c.OnError(func(r *colly.Response, err error) {
		//logger.Error("OnError", zap.String("error", err.Error()))
	})

	err = c.Visit(enterUrl)
	return dataResSlice, err

}

// toParseUl 解析ul标签
func toParseLi(liElement *colly.HTMLElement) projectInfo {
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
	projectDesc = strings.SplitN(projectDesc, "\n", 2)[0]
	projectInfo := projectInfo{
		ProjectName: projectName,
		Addr:        addr,
		ProjectDesc: projectDesc,
	}
	return projectInfo
}
