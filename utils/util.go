package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

func Validate(v interface{}) ([]string, error) {
	validate := validator.New()
	err := validate.Struct(v)
	if err == nil {
		return nil, err
	}

	v, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil, errors.New("invalid type")
	}

	messages := []string{}
	for _, err := range err.(validator.ValidationErrors) {
		messages = append(messages, fmt.Sprintf("%s is %s", strcase.ToSnake(err.Field()), err.Tag()))
	}

	return messages, err
}
