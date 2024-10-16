package pkg

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(data interface{}) error {
	validate := validator.New()

	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	validationErros := err.(validator.ValidationErrors)
	validationError := validationErros[0]

	switch validationError.Tag() {
	case "required":
		return errors.New(validationError.StructField() + " is required")
	case "max":
		return errors.New(validationError.StructField() + " is required with max: " + validationError.Param())
	case "min":
		return errors.New(validationError.StructField() + " is required with min: " + validationError.Param())
	case "email":
		return errors.New(validationError.StructField() + " is invalid.")
	}

	return nil
}
