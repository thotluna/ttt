package validator

// MockValidator es un mock de la interfaz Validator para usar en pruebas
type MockValidator struct {
	// ValidateFunc almacena la función que será llamada por el método Validate
	// Si es nil, Validate retornará nil
	ValidateFunc func(interface{}) error

	// NextValidator almacena el siguiente validador en la cadena
	// Si es nil, Next() retornará nil
	NextValidator Validator
}

// Validate implementa la interfaz Validator
// Si ValidateFunc no es nil, la ejecuta y retorna su resultado
// Si ValidateFunc es nil, retorna nil
func (m *MockValidator) Validate(value interface{}) error {
	if m.ValidateFunc != nil {
		return m.ValidateFunc(value)
	}
	return nil
}

// Next implementa la interfaz Validator
// Retorna el validador configurado en NextValidator
func (m *MockValidator) Next() Validator {
	return m.NextValidator
}

// NewMockValidator crea una nueva instancia de MockValidator
// con la función de validación y el siguiente validador opcionales
func NewMockValidator(validateFunc func(interface{}) error, next Validator) *MockValidator {
	return &MockValidator{
		ValidateFunc: validateFunc,
		NextValidator: next,
	}
}
