/*
 * @date: 2021/12/15
 * @desc: ...
 */

package global

import (
	"github.com/daniuEvan/mygithub/config"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{} // 全局配置
	Logger       *zap.Logger                                   // 全局logger
	Trans        ut.Translator                                 // 错误表单校验
	OrmDB        *gorm.DB                                      // 数据库连接
)

var (
	SpiderLogger *zap.Logger // 爬虫logger
)
