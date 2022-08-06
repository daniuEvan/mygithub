/**
 * @date: 2022/7/31
 * @desc:
 */

package awesomeSpider

import (
	"github.com/daniuEvan/mygithub/global"
	"github.com/daniuEvan/mygithub/internal/model/spiderModel"
	"gorm.io/gorm/clause"
)

func toDatabase(awesome *spiderModel.Awesome) (err error) {
	db := global.OrmDB
	err = db.Clauses(clause.OnConflict{UpdateAll: true}).Create(awesome).Error
	return err
}
