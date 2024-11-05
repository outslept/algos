package lipogram

import (
        "testing"
)

func TestIsLipogram(t *testing.T) {
        tests := []struct {
                s              string
                excludedLetter rune
                expected       bool
        }{
                {"the quick brown fox", 'a', true},
                {"the quick brown fox jumps over the lazy dog", 'e', false},
                {"the quick brown fox jumps over the lazy dog", 'x', true},
                {"", 'a', true}, // Edge case: empty string
                {"hello", 'h', false},
                {"hello", 'z', true},
        }

        for _, test := range tests {
                t.Run(test.s, func(t *testing.T) {
                        result := isLipogram(test.s, test.excludedLetter)
                        if result != test.expected {
                                t.Errorf("Expected isLipogram(%q, %q) = %v, but got %v", test.s, string(test.excludedLetter), test.expected, result)
                        }
                })
        }
}
