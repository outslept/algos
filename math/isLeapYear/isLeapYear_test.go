package leapyear

import (
        "testing"
				"fmt"
)

func TestIsLeapYear(t *testing.T) {
        tests := []struct {
                year     int
                expected bool
        }{
                {1996, true},  // Divisible by 4, not 100
                {1900, false}, // Divisible by 100, not 400
                {2000, true},  // Divisible by 400
                {2021, false}, // Not divisible by 4
                {2024, true},  // Upcoming leap year
                {2100, false}, // Divisible by 100, not 400
                {2400, true},  // Divisible by 400
        }

        for _, test := range tests {
                t.Run(fmt.Sprintf("isLeapYear(%d)", test.year), func(t *testing.T) {
                        result := isLeapYear(test.year)
                        if result != test.expected {
                                t.Errorf("Expected isLeapYear(%d) = %v, but got %v", test.year, test.expected, result)
                        }
                })
        }
}
