package validator

import (
	"reflect"
	"strings"

	validation "github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validation.Validate
}

func NewValidator() *CustomValidator {
	v := validation.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})
	return &CustomValidator{
		validator: v,
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		var errVal ErrorValidation
		for _, e := range err.(validation.ValidationErrors) {
			fe := e.Tag()
			if param := e.Param(); len(param) >= 0 {
				fe = fe + ":" + param
			}
			e := FieldError{
				Field: e.Field(),
				Error: fe,
			}
			errVal = append(errVal, e)
		}

		return &errVal
	}

	return nil
}

func (cv *CustomValidator) ValidateStruct(i interface{}) error {
	return cv.validator.Struct(i)
}
