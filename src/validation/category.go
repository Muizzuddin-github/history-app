package validation

import (
	"crud/src/requestbody"

	"github.com/go-playground/validator/v10"
)

func Category(data *requestbody.AddCategory) []string{ 
	val := validator.New(validator.WithRequiredStructEnabled())
	err := val.Struct(data)
	result := []string{}
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			result = append(result, err.Error())
		}
	}
	return result
}
func CategoryUp(data *requestbody.UpdateCategory) []string{ 
	val := validator.New(validator.WithRequiredStructEnabled())
	err := val.Struct(data)
	result := []string{}
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			result = append(result, err.Error())
		}
	}
	return result
}
