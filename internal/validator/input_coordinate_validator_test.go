package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInputCoordinateValidator(t *testing.T) {
	tests := []struct {
		name string
		next Validator
	}{
		{
			name: "with next validator",
			next: NewMockValidator(nil, nil),
		},
		{
			name: "without next validator",
			next: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewInputCoordinateValidator(tt.next)
			assert.NotNil(t, v)
		})
	}
}

func TestInputCoordinateValidator_Validate(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		wantErr bool
		errCode string
	}{
		{
			name:    "valid coordinate",
			input:   "1.2",
			wantErr: false,
		},
		{
			name:    "valid coordinate with spaces",
			input:   " 1 . 2 ",
			wantErr: false,
		},
		{
			name:    "nil input",
			input:   nil,
			wantErr: true,
			errCode: CodeInvalidInput,
		},
		{
			name:    "not a string",
			input:   123,
			wantErr: true,
			errCode: CodeInvalidInput,
		},
		{
			name:    "invalid format - missing dot",
			input:   "12",
			wantErr: true,
			errCode: CodeInvalidFormat,
		},
		{
			name:    "invalid format - multiple dots",
			input:   "1.2.3",
			wantErr: true,
			errCode: CodeInvalidFormat,
		},
		{
			name:    "invalid format - non-numeric characters",
			input:   "a.b",
			wantErr: true,
			errCode: CodeInvalidFormat,
		},
		{
			name:    "invalid format - dot at start",
			input:   ".2",
			wantErr: true,
			errCode: CodeInvalidFormat,
		},
		{
			name:    "invalid format - dot at end",
			input:   "1.",
			wantErr: true,
			errCode: CodeInvalidFormat,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewInputCoordinateValidator(nil)
			err := v.Validate(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				errVal, ok := err.(*Error)
				assert.True(t, ok)
				assert.Equal(t, tt.errCode, errVal.Code)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestInputCoordinateValidator_WithNextValidator(t *testing.T) {
	mockNext := NewMockValidator(
		func(interface{}) error {
			return NewError("NEXT_ERROR", "error from next validator")
		},
		nil,
	)

	v := NewInputCoordinateValidator(mockNext)
	err := v.Validate("1.2")

	errVal, ok := err.(*Error)
	assert.True(t, ok)
	assert.Equal(t, "NEXT_ERROR", errVal.Code)
}


