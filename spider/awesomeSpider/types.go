/**
 * @date: 2022/7/27
 * @desc:
 */

package awesomeSpider

type projectInfo struct {
	ProjectName    string
	Addr           string
	Author         string
	StarCount      int
	LastCommitTime string
	ProjectDesc    string
}

type DataRes struct {
	DevLangCategory     string //开发语言类别
	PurposeCategory     string //用途分类
	PurposeCategoryDesc string //用途描述
	ProjectInfos        []projectInfo
}
