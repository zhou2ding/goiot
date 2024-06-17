package trans

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func GetTrans() ut.Translator {
	return trans
}

func InitTranslatorOfValidator(local string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT, enT)

		if trans, ok = uni.GetTranslator(local); !ok {
			return errors.New("invalid local language")
		}
		// 注册翻译器
		switch local {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
			_ = v.RegisterTranslation("password", trans, registerEnFunc, translateFunc)
		case "zh":
			err = chTranslations.RegisterDefaultTranslations(v, trans)
			_ = v.RegisterTranslation("password", trans, registerZhFunc, translateFunc)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
			_ = v.RegisterTranslation("password", trans, registerEnFunc, translateFunc)
		}
		return
	}
	return
}

func registerEnFunc(ut ut.Translator) error {
	return ut.Add("password", "{0} must include at least 8 characters, and include letter, number and special character", true)
}

func registerZhFunc(ut ut.Translator) error {
	return ut.Add("password", "{0}必须包含至少8个字符，且包含字母、数字、特殊字符", true)
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("password", fe.Field())
	return t
}
