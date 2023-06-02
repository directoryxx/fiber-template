package validation

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ErrorValidationResponse - Standarize the response for validation
type ErrorValidationResponse struct {
	FailedField string
	Tag         string
	Value       string
	Message     string
}

// ValidateStruct - Validate Input for all usecase
func ValidateStruct(class interface{}) []*ErrorValidationResponse {
	var errors []*ErrorValidationResponse

	err := validate.Struct(class)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorValidationResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
