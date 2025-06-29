package validator

type MockValidator struct {
	ValidateFunc func(interface{}) error
	NextValidator Validator
}

func (m *MockValidator) Validate(value interface{}) error {
	if m.ValidateFunc != nil {
		return m.ValidateFunc(value)
	}
	return nil
}

func (m *MockValidator) Next() Validator {
	return m.NextValidator
}

func NewMockValidator(validateFunc func(interface{}) error, next Validator) *MockValidator {
	return &MockValidator{
		ValidateFunc: validateFunc,
		NextValidator: next,
	}
}
