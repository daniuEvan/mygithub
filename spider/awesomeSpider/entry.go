/**
 * @date: 2022/7/28
 * @desc:
 */

package awesomeSpider

import (
	"errors"
	"go.uber.org/zap"
)

func AwesomeSpider(langCategory string, logger *zap.Logger) (dataResSlice []DataRes, err error) {
	addr, _ := awesomeDatasource[langCategory]
	switch langCategory {
	case "awesome-python":
		dataResSlice, err = awesomePythonSpider(addr, logger)
	case "awesome-go":
		dataResSlice, err = awesomeGoSpider(addr, logger)
	default:
		return nil, errors.New("爬虫类型错误")
	}
	return dataResSlice, err
}
