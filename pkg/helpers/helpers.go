package helpers

import "github.com/go-playground/validator/v10"

func FormValidationError(err error) []string{
	var errors []string
	// switch data type from error to validator error
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}