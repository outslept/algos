package heterogram

import (
        "testing"
)

func TestIsHeterogram(t *testing.T) {
        tests := []struct {
                s        string
                expected bool
        }{
                {"the quick brown", true},
                {"jumped over", true},
                {"lazy dog", true},
                {"hello", false},
                {"", true}, // Edge case: empty string
                {"a", true},
                {"aa", false},
                {"I am", false}, // Duplicate 'a'
        }

        for _, test := range tests {
                t.Run(test.s, func(t *testing.T) {
                        result := isHeterogram(test.s)
                        if result != test.expected {
                                t.Errorf("Expected isHeterogram(%q) = %v, but got %v", test.s, test.expected, result)
                        }
                })
        }
}
