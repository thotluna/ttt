package validator

import "regexp"

var coordRegexp = regexp.MustCompile(`^\s*\d+\s*\.\s*\d+\s*$`)

type InputCoordinateValidator struct {
	ValidatorBase
}

func NewInputCoordinateValidator(next Validator) *InputCoordinateValidator {
	return &InputCoordinateValidator{
		ValidatorBase: ValidatorBase{
			other: next,
		},
	}
}

func (v *InputCoordinateValidator) Validate(value interface{}) error {
	if value == nil {
		return NewError(CodeInvalidInput, MsgExpectedString)
	}

	str, ok := value.(string)
	if !ok {
		return NewError(CodeInvalidInput, MsgExpectedString)
	}

	if !coordRegexp.MatchString(str) {
		return NewError(CodeInvalidFormat, MsgInvalidFormat)
	}

	if v.other != nil {
		return v.other.Validate(value)
	}

	return nil
}
