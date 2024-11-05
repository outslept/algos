package perfectnumber

import (
        "testing"
)

func TestIsPerfectNumber(t *testing.T) {
        tests := []struct {
                n        int
                expected bool
        }{
                {6, true},   // Perfect number
                {28, true},  // Perfect number
                {496, true}, // Perfect number
                {12, false}, // Not a perfect number
                {1, false},  // Not a perfect number
                {0, false},  // Not a perfect number
                {8128, true}, // Perfect number
        }

        for _, test := range tests {
                t.Run(fmt.Sprintf("isPerfectNumber(%d)", test.n), func(t *testing.T) {
                        result := isPerfectNumber(test.n)
                        if result != test.expected {
                                t.Errorf("Expected isPerfectNumber(%d) = %v, but got %v", test.n, test.expected, result)
                        }
                })
        }
}
