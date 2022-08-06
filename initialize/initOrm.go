/**
 * @date: 2022/7/31
 * @desc:
 */

package initialize

import (
	"errors"
	"github.com/daniuEvan/mygithub/global"
	"github.com/daniuEvan/mygithub/internal/model"
	"gorm.io/gorm"
)

// InitOrm 初始化数据库
func InitOrm() (err error) {
	var db *gorm.DB
	switch global.ServerConfig.DatabaseInfo.DBType {
	case "mysql":
		db, err = initOrmMysql()
		if err != nil {
			return err
		}

	case "postgres":
		db, err = initOrmPostgres()
		if err != nil {
			return err
		}
	default:
		return errors.New("数据库类型不匹配, 请检查配置文件")
	}

	// 数据库表迁移
	err = db.AutoMigrate(model.ModelsArr...)
	if err != nil {
		return err
	}
	global.OrmDB = db
	return err
}
