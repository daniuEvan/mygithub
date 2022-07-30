/**
 * @date: 2022/7/28
 * @desc: 爬虫周期配置信息
 */

package config

type Spider struct {
	AwesomeGo     string `mapstructure:"awesome-go" json:"awesome-go"`
	AwesomePython string `mapstructure:"awesome-python" json:"awesome-python"`
}
