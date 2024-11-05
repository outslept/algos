package tautogram

import (
        "testing"
)

func TestIsTautogram(t *testing.T) {
        tests := []struct {
                s        string
                expected bool
        }{
                {"She sells sea shells", true},
                {"Peter Piper picked a peck", true},
                {"The quick brown fox", false},
                {"", true}, // Edge case: empty string
                {"A", true}, // Single word
                {"A apple", true},
                {"an apple", true},
                {"An apple a day", false},
        }

        for _, test := range tests {
                t.Run(test.s, func(t *testing.T) {
                        result := isTautogram(test.s)
                        if result != test.expected {
                                t.Errorf("Expected isTautogram(%q) = %v, but got %v", test.s, test.expected, result)
                        }
                })
        }
}
