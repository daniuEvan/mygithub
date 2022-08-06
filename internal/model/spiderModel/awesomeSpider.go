/**
 * @date: 2022/7/31
 * @desc:
 */

package spiderModel

import "time"

type Awesome struct {
	//gorm.Model
	DevLangCategory     string `gorm:"primaryKey;type:varchar(100);column:dev_lang_category;NOT NULL;comment:开发语言类型"`
	PurposeCategory     string `gorm:"primaryKey;type:varchar(100);column:purpose_category;NOT NULL;comment:开源库的作用"`
	PurposeCategoryDesc string `gorm:"column:purpose_category_desc;comment:开源库作用描述"`
	ProjectName         string `gorm:"primaryKey;type:varchar(100);column:project_name;NOT NULL;comment:开源库名称"`
	Addr                string `gorm:"type:varchar(100);column:addr;NOT NULL;comment:github地址"`
	IsGithub            int    `gorm:"type:int;column:addr;NOT NULL;comment:是否为github托管,0-no,1-yes"`
	Author              string `gorm:"type:varchar(100);column:author;NOT NULL;comment:作者"`
	StarCount           string `gorm:"type:varchar(100);column:star_count;NOT NULL;comment:星星数量"`
	LastCommitTime      string `gorm:"type:varchar(100);column:last_commit_time;NOT NULL;comment:最后提交时间"`
	ProjectDesc         string `gorm:"column:project_desc;comment:开源库项目描述"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (Awesome) TableName() string {
	return "awesome"
}
