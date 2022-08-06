/**
 * @date: 2022/7/31
 * @desc:
 */

package initialize

import (
	"fmt"
	"github.com/daniuEvan/mygithub/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func initOrmMysql() (db *gorm.DB, err error) {
	mysqlInfo := global.ServerConfig.DatabaseInfo.MysqlInfo
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.Username,
		mysqlInfo.Password,
		mysqlInfo.Host,
		mysqlInfo.Port,
		mysqlInfo.DBName,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%s", global.ServerConfig.DatabaseInfo.TablePrefix),
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(mysqlInfo.MaxIdleConns)
	sqlDB.SetMaxOpenConns(mysqlInfo.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlInfo.ConnMaxLifetime) * time.Second)
	return
}
