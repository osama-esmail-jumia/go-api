package validation

import "github.com/go-playground/validator/v10"

func Validate(v interface{}) error {
	var validate *validator.Validate
	validate = validator.New()
	err := validate.Struct(v)
	return err
}
