/*
 * @date: 2021/12/15
 * @desc: ...
 */

package initialize

import (
	"fmt"
	"github.com/daniuEvan/mygithub/global"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

// initTrans 翻译
func initTrans(locale string) (err error) {
	//修改gin框架中的validator引擎属性, 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		// 第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}
		switch locale {
		case "en":
			if err := enTranslations.RegisterDefaultTranslations(v, global.Trans); err != nil {
				panic(err)
			}
		case "zh":
			if err := zhTranslations.RegisterDefaultTranslations(v, global.Trans); err != nil {
				panic(err)
			}
		default:
			if err := enTranslations.RegisterDefaultTranslations(v, global.Trans); err != nil {
				panic(err)
			}
		}
		return
	}
	return
}
