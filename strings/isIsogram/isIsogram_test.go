package isogram

import (
        "testing"
)

func TestIsIsogram(t *testing.T) {
        tests := []struct {
                s        string
                expected bool
        }{
                {"isogram", true},
                {"background", true},
                {"downstream", true},
                {"six-year-old", true}, // Hyphen ignored
                {"hello", false},
                {"", true}, // Edge case: empty string
                {"a", true},
                {"aa", false},
                {"Alphabet", false}, // Case insensitive check
        }

        for _, test := range tests {
                t.Run(test.s, func(t *testing.T) {
                        result := isIsogram(test.s)
                        if result != test.expected {
                                t.Errorf("Expected isIsogram(%q) = %v, but got %v", test.s, test.expected, result)
                        }
                })
        }
}
