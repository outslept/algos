package twopointers

import (
    "reflect"
    "testing"
)

func TestTwoPointers(t *testing.T) {
    tests := []struct {
        name     string
        nums     []int
        target   int
        expected [2]int
    }{
        {
            name:     "basic case",
            nums:     []int{2, 3, 4, 7, 11, 15},
            target:   9,
            expected: [2]int{0, 3},
        },
        {
            name:     "no solution",
            nums:     []int{1, 2, 3, 4},
            target:   10,
            expected: [2]int{-1, -1},
        },
        {
            name:     "numbers at edges",
            nums:     []int{1, 2, 3, 4, 5},
            target:   6,
            expected: [2]int{0, 4},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := TwoPointers(tt.nums, tt.target)
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("TwoPointers() = %v, want %v", result, tt.expected)
            }
        })
    }
}
