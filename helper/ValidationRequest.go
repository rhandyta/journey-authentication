package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func RequestValidationHelper(paramStruct interface{}) string {
	var validationMessage string
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(paramStruct)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return ValidationErrorToText(e)
			break
		}
	}

	return validationMessage
}

func ValidationErrorToText(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required !", e.Field())
	case "required_if":
		return fmt.Sprintf("%s is required !", e.Field())
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s !", e.Field(), e.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s !", e.Field(), e.Param())
	case "len":
		return fmt.Sprintf("%s must be %s characters long !", e.Field(), e.Param())

	}
	return fmt.Sprintf("%s is not valid", e.Field())
}
