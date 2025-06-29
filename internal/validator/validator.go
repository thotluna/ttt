package validator

import "fmt"

// ValidationError representa un error de validación
type ValidationError struct {
	Code    string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// NewValidationError crea un nuevo error de validación
func NewValidationError(code, format string, args ...interface{}) *ValidationError {
	return &ValidationError{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}

type Validator interface {
	Validate(value interface{}) error
	Next() Validator
}

type ValidatorBase struct {
	other Validator
}
