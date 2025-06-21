package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Valid *validator.Validate
}

func InitValidator() *Validator {
	validator := validator.New()
	validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "_" {
			return ""
		}
		return name
	})
	return &Validator{
		Valid: validator,
	}
}

func (v *Validator) ValidateStruct(s interface{}) error {
	if err := v.Valid.Struct(s); err != nil {
		var validateErr validator.ValidationErrors
		if errors.As(err, &validateErr) {
			return convertToUserError(validateErr)
		}
	}
	return nil
}

func convertToUserError(err error) error {
	switch t := err.(type) {
	case validator.ValidationErrors:
		var validationErr string
		for _, v := range t {
			validationErr += fmt.Sprintf("%s - %s ", v.Field(), convertTag(v.Tag()))
		}
		validationErr = "Validation failed for the following field(s): " + validationErr[:len(validationErr)-2]
		return errors.New(validationErr)
	case *validator.InvalidValidationError:
		return fmt.Errorf("cannot validate request with body: %s", t.Type)
	default:
		return t

	}
}

func convertTag(tag string) string {
	switch tag {
	case "required":
		return "is required"
	default:
		return tag
	}
}
