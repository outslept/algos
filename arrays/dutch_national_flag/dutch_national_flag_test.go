package dutch

import (
	"reflect"
	"testing"
)

func TestDutchFlag(t *testing.T) {
    tests := []struct {
        name     string
        input    []int
        expected []int
    }{
        {
            name:     "basic case",
            input:    []int{2, 0, 1, 2, 1, 0},
            expected: []int{0, 0, 1, 1, 2, 2},
        },
        {
            name:     "already sorted",
            input:    []int{0, 0, 1, 1, 2, 2},
            expected: []int{0, 0, 1, 1, 2, 2},
        },
        {
            name:     "all same numbers",
            input:    []int{1, 1, 1},
            expected: []int{1, 1, 1},
        },
        {
            name:     "empty array",
            input:    []int{},
            expected: []int{},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            input := make([]int, len(tt.input))
            copy(input, tt.input)
            
            DutchFlag(input)
            
            if !reflect.DeepEqual(input, tt.expected) {
                t.Errorf("DutchFlag() = %v, want %v", input, tt.expected)
            }
        })
    }
}

func BenchmarkDutchFlag(b *testing.B) {
    nums := []int{2, 0, 1, 2, 1, 0}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        input := make([]int, len(nums))
        copy(input, nums)
        DutchFlag(input)
    }
}
