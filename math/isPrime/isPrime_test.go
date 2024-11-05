package prime

import (
        "testing"
				"fmt"
)

func TestIsPrime(t *testing.T) {
        tests := []struct {
                n        int
                expected bool
        }{
                {2, true},    // Smallest prime number
                {3, true},    // Prime number
                {4, false},   // Not a prime number
                {5, true},    // Prime number
                {9, false},   // Not a prime number
                {11, true},   // Prime number
                {1, false},   // Not a prime number
                {0, false},   // Not a prime number
                {-1, false},  // Not a prime number
                {17, true},   // Prime number
                {19, true},   // Prime number
                {20, false},  // Not a prime number
                {23, true},   // Prime number
                {29, true},   // Prime number
        }

        for _, test := range tests {
                t.Run(fmt.Sprintf("isPrime(%d)", test.n), func(t *testing.T) {
                        result := isPrime(test.n)
                        if result != test.expected {
                                t.Errorf("Expected isPrime(%d) = %v, but got %v", test.n, test.expected, result)
                        }
                })
        }
}
