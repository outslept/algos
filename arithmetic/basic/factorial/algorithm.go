package factorial

import (
	"errors"
	"fmt"
)

// Calculate computes the factorial of n iteratively.
// Returns error for negative numbers or when result exceeds int64.
func Calculate(n int64) (int64, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}

	// Handle base cases
	if n == 0 || n == 1 {
		return 1, nil
	}

	// Maximum value that can be computed without overflow in int64
	if n > 20 {
		return 0, fmt.Errorf("factorial of %d is too large for int64", n)
	}

	result := int64(1)
	for i := int64(2); i <= n; i++ {
		result *= i
		// Check for overflow (though redundant with n>20 check above)
		if result < 0 {
			return 0, fmt.Errorf("factorial computation overflow at n=%d", n)
		}
	}

	return result, nil
}

// CalculateRecursive computes the factorial of n recursively.
// Returns error for negative numbers or when result exceeds int64.
func CalculateRecursive(n int64) (int64, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}

	// Handle base cases
	if n == 0 || n == 1 {
		return 1, nil
	}

	// Check for potential overflow
	if n > 20 {
		return 0, fmt.Errorf("factorial of %d is too large for int64", n)
	}

	// Recursive calculation
	prev, err := CalculateRecursive(n - 1)
	if err != nil {
		return 0, err
	}

	result := n * prev
	// Check for overflow (though redundant with n>20 check above)
	if result < 0 {
		return 0, fmt.Errorf("factorial computation overflow at n=%d", n)
	}

	return result, nil
}
