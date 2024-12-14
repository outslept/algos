package factorial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFactorial_Calculate verifies factorial calculation:
// - Base cases (0!, 1!)
// - Small numbers
// - Error cases (negative numbers)
func TestFactorial_Calculate(t *testing.T) {
	tests := []struct {
		name          string
		input         int64
		expected      int64
		expectedError bool
	}{
		{
			name:          "Factorial of 0",
			input:         0,
			expected:      1,
			expectedError: false,
		},
		{
			name:          "Factorial of 1",
			input:         1,
			expected:      1,
			expectedError: false,
		},
		{
			name:          "Factorial of 5",
			input:         5,
			expected:      120,
			expectedError: false,
		},
		{
			name:          "Factorial of 10",
			input:         10,
			expected:      3628800,
			expectedError: false,
		},
		{
			name:          "Negative number",
			input:         -1,
			expected:      0,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Calculate(tt.input)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

// TestFactorial_RecursiveCalculate verifies recursive factorial implementation:
// - Compares results with iterative version
// - Validates error handling
func TestFactorial_RecursiveCalculate(t *testing.T) {
	tests := []struct {
		name          string
		input         int64
		expected      int64
		expectedError bool
	}{
		{
			name:          "Recursive factorial of 0",
			input:         0,
			expected:      1,
			expectedError: false,
		},
		{
			name:          "Recursive factorial of 6",
			input:         6,
			expected:      720,
			expectedError: false,
		},
		{
			name:          "Negative number",
			input:         -5,
			expected:      0,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CalculateRecursive(tt.input)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)

				// Compare with iterative version
				iterResult, iterErr := Calculate(tt.input)
				assert.NoError(t, iterErr)
				assert.Equal(t, iterResult, result)
			}
		})
	}
}

// TestFactorial_EdgeCases verifies handling of edge cases:
// - Large numbers (overflow check)
// - Zero and one as special cases
func TestFactorial_EdgeCases(t *testing.T) {
	t.Run("Large number overflow check", func(t *testing.T) {
		_, err := Calculate(21) // 21! exceeds int64
		assert.Error(t, err)
	})

	t.Run("Special cases", func(t *testing.T) {
		// 0! = 1
		result, err := Calculate(0)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), result)

		// 1! = 1
		result, err = Calculate(1)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), result)
	})
}

// TestFactorial_Validation verifies input validation:
// - Negative numbers
// - Boundary conditions
func TestFactorial_Validation(t *testing.T) {
	t.Run("Negative numbers", func(t *testing.T) {
		_, err := Calculate(-10)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "negative")
	})

	t.Run("Maximum allowed input", func(t *testing.T) {
		result, err := Calculate(20) // 20! is the maximum that fits in int64
		assert.NoError(t, err)
		assert.Equal(t, int64(2432902008176640000), result)
	})
}
