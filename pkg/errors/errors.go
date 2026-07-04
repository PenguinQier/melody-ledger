package errors

import (
	"errors"
	"github.com/PenguinQier/melody-ledger/internal/providers/validation"
	"github.com/go-playground/validator/v10"
	"strings"
)

var errorsList = make(map[string]string)

func Init() {
	errorsList = map[string]string{}
}

func SetFormErrors(err error) {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			Add(fieldError.Field(), GetErrorMsg(fieldError.Tag()))
		}
	}
}

func Add(key string, value string) {
	errorsList[strings.ToLower(key)] = value
}

func GetErrorMsg(tag string) string {
	return validation.ErrorMessage()[tag]
}

func Get() map[string]string {
	return errorsList
}
