package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
    tests := []struct {
        name     string
        nums     []int
        target   int
        expected [2]int
    }{
        {
            name:     "basic case",
            nums:     []int{2, 7, 11, 15},
            target:   9,
            expected: [2]int{0, 1},
        },
        {
            name:     "no solution",
            nums:     []int{1, 2, 3},
            target:   7,
            expected: [2]int{-1, -1},
        },
        {
            name:     "duplicate numbers",
            nums:     []int{3, 3},
            target:   6,
            expected: [2]int{0, 1},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := TwoSum(tt.nums, tt.target)
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("TwoSum() = %v, want %v", result, tt.expected)
            }
        })
    }
}

func BenchmarkTwoSum(b *testing.B) {
    nums := []int{2, 7, 11, 15}
    target := 9
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        TwoSum(nums, target)
    }
}
