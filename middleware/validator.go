package middleware

import "github.com/go-playground/validator/v10"

// CustomValidator holds the validator instance.
type CustomValidator struct {
	validator *validator.Validate
}

// NewCustomValidator creates a new instance of CustomValidator.
func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

// Validate validates the given struct using the validator instance.
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return err
	}
	return nil
}
