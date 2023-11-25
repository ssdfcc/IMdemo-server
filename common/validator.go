package common

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"regexp"
)

var rules = map[string]validator.Func{
	"regex": hasValidRegex,
}

// 注册自定义校验规则
func RegisterValidation(validate *validator.Validate, trans ut.Translator) {
	for key, value := range rules {
		_ = validate.RegisterValidation(key, value)
		_ = validate.RegisterTranslation(key, trans, func(ut ut.Translator) error {
			return ut.Add(key, "{0}参数错误", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(key, fe.Field())
			return t
		})
	}
}

// 正则校验
func hasValidRegex(fl validator.FieldLevel) bool {
	regexPattern := fl.Param()
	re, err := regexp.Compile(regexPattern)
	if err != nil {
		panic(fmt.Sprintf("Bad field type %T", fl.Field().Interface()))
		return false
	}
	fieldValue := fl.Field().String()
	return re.MatchString(fieldValue)
}
