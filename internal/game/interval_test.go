package game

import "testing"

func TestGameInterval(t *testing.T) {
	interval := NewGameInterval()

	tests := []struct {
		name     string
		value    int
		expected bool
	}{
		{
			name:     "min value",
			value:    0,
			expected: true,
		},
		{
			name:     "max value",
			value:    2,
			expected: true,
		},
		{
			name:     "middle value",
			value:    1,
			expected: true,
		},
		{
			name:     "below min",
			value:    -1,
			expected: false,
		},
		{
			name:     "above max",
			value:    3,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := interval.Contains(tt.value)
			if result != tt.expected {
				t.Errorf("Expected %v for value %d, got %v", tt.expected, tt.value, result)
			}

			if interval.Min() != 0 {
				t.Error("Expected min to be 0")
			}

			if interval.Max() != 2 {
				t.Error("Expected max to be 2")
			}
		})
	}
}
