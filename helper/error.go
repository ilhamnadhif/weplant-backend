package helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err.Error())
	}
}


func IfValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
		errors = append(errors, errorMessage)
	}
	return errors
}
