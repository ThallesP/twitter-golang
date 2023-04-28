package utils

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	english "github.com/go-playground/validator/v10/translations/en"
	"github.com/thallesp/twitter-golang/usecase"
)

var validate = validator.New()
var uni = ut.New(en.New(), en.New())

func ValidateStruct(s interface{}) []*usecase.Exception {
	trans, _ := uni.GetTranslator("en")

	english.RegisterDefaultTranslations(validate, trans)

	var errors []*usecase.Exception

	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, usecase.NewException(err.Translate(trans), 400, "bad_request"))
		}
	}
	return errors
}
