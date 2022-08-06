/**
 * @date: 2022/8/6
 * @desc:
 */

package getStarCount

import (
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

//
// GetStarCount
// @Description: 获取星星数量
// @param githubAddr:
// @param logger:
// @return startCount:
// @return lastCommit:
// @return err:
//
func GetStarCount(githubAddr string, logger *zap.Logger) (startCount, lastCommit string, err error) {
	//logger := global.Logger
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	c.OnHTML("span[id=repo-stars-counter-star]", func(element *colly.HTMLElement) {
		startCount = element.Text
	})
	c.OnHTML("a[class='Link--secondary ml-2'] > relative-time[class='no-wrap']", func(element *colly.HTMLElement) {
		lastCommit = element.Text
		//fmt.Println("commit.Text:", element.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		logger.Info("StarCountOnRequest :", zap.String("url", r.URL.String()))
	})

	c.OnResponse(func(r *colly.Response) {
		//time.Sleep(500 * time.Millisecond)
		//logger.Info("OnResponse :", zap.String("url", fmt.Sprintf("Response %s: %d bytes", r.Request.URL, len(r.Body))))

	})

	c.OnError(func(r *colly.Response, err error) {
		//logger.Error("OnError", zap.String("error", err.Error()))
	})

	err = c.Visit(githubAddr)
	return startCount, lastCommit, err
}
