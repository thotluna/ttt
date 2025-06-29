package validator

import "testing"

func TestInterval_Contains(t *testing.T) {
	tests := []struct {
		name     string
		min      int
		max      int
		value    int
		expected bool
	}{
		{
			name:     "value equals min",
			min:      1,
			max:      10,
			value:    1,
			expected: true,
		},
		{
			name:     "value equals max",
			min:      1,
			max:      10,
			value:    10,
			expected: true,
		},
		{
			name:     "value in middle",
			min:      1,
			max:      10,
			value:    5,
			expected: true,
		},
		{
			name:     "value below min",
			min:      1,
			max:      10,
			value:    0,
			expected: false,
		},
		{
			name:     "value above max",
			min:      1,
			max:      10,
			value:    11,
			expected: false,
		},
		{
			name:     "negative range",
			min:      -5,
			max:      5,
			value:    0,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interval := NewInterval(tt.min, tt.max)
			result := interval.Contains(tt.value)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v for value %d in range [%d, %d]",
					tt.expected, result, tt.value, tt.min, tt.max)
			}

			// Test getters
			if interval.Min() != tt.min {
				t.Errorf("Expected min %d, got %d", tt.min, interval.Min())
			}
			if interval.Max() != tt.max {
				t.Errorf("Expected max %d, got %d", tt.max, interval.Max())
			}
		})
	}
}
