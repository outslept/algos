package pangram

import (
        "testing"
)

func TestIsPangram(t *testing.T) {
        tests := []struct {
                s        string
                expected bool
        }{
                {"The quick brown fox jumps over the lazy dog", true},
                {"Pack my box with five dozen liquor jugs", true},
                {"Sphinx of black quartz, judge my vow", true},
                {"Hello world", false},
                {"", false}, // Edge case: empty string
                {"abcdefghijklmnopqrstuvwxyz", true},
                {"a quick movement of the enemy will jeopardize five gunboats", false}, // Missing 'x'
        }

        for _, test := range tests {
                t.Run(test.s, func(t *testing.T) {
                        result := isPangram(test.s)
                        if result != test.expected {
                                t.Errorf("Expected isPangram(%q) = %v, but got %v", test.s, test.expected, result)
                        }
                })
        }
}
