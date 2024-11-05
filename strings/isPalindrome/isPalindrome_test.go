package palindrome

import (
        "testing"
)

func TestIsPalindrome(t *testing.T) {
        tests := []struct {
                s        string
                expected bool
        }{
                {"racecar", true},
                {"level", true},
                {"hello", false},
                {"", true}, // Edge case: empty string
                {"a", true},
                {"ab", false},
                {"abba", true},
                {"abcdcba", true},
                {"abcd", false},
        }

        for _, test := range tests {
                t.Run(test.s, func(t *testing.T) {
                        result := isPalindrome(test.s)
                        if result != test.expected {
                                t.Errorf("Expected isPalindrome(%q) = %v, but got %v", test.s, test.expected, result)
                        }
                })
        }
}
