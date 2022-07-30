/*
 * @date: 2021/12/15
 * @desc: ...
 */

package initialize

import (
	"github.com/daniuEvan/mygithub/common/customValidator"
	"github.com/daniuEvan/mygithub/global"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// initCustomValidator 初始化自定义校验器
func initCustomValidator() (err error) {
	err = registerMobileValidator()
	if err != nil {
		return err
	}
	return nil
}

// registerMobileValidator 手机号码校验器
func registerMobileValidator() (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err = v.RegisterValidation("mobile", customValidator.ValidateMobile)
		if err != nil {
			return err
		}
		err = v.RegisterTranslation(
			"mobile",
			global.Trans,
			func(ut ut.Translator) error {
				return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
			},
			func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("mobile", fe.Field())
				return t
			})
		if err != nil {
			return err
		}
	}
	return nil
}
