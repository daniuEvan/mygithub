/**
 * @date: 2022/7/31
 * @desc:
 */

package initialize

import (
	"fmt"
	"github.com/daniuEvan/mygithub/global"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func initOrmPostgres() (db *gorm.DB, err error) {
	postgresInfo := global.ServerConfig.DatabaseInfo.PostgresInfo
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s  port=%d  dbname=%s search_path=%s sslmode=disable TimeZone=Asia/Shanghai",
		postgresInfo.Username,
		postgresInfo.Password,
		postgresInfo.Host,
		postgresInfo.Port,
		postgresInfo.DBName,
		postgresInfo.Schema,
	)
	tablePrefix := global.ServerConfig.DatabaseInfo.TablePrefix
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%s.%s", postgresInfo.Schema, tablePrefix), // pg 加上schema后 表前缀失效
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(postgresInfo.MaxIdleConns)
	sqlDB.SetMaxOpenConns(postgresInfo.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(postgresInfo.ConnMaxLifetime) * time.Second)
	return

}
