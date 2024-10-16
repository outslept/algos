package oddevensort

import (
	"reflect"
	"testing"
)

func TestOddEvenSort(t *testing.T) {
        tests := []struct {
                name     string
                input    []int
                expected []int
        }{
                {
                        name:     "Empty slice",
                        input:    []int{},
                        expected: []int{},
                },
                {
                        name:     "Single element",
                        input:    []int{1},
                        expected: []int{1},
                },
                {
                        name:     "Already sorted",
                        input:    []int{1, 2, 3, 4, 5},
                        expected: []int{1, 2, 3, 4, 5},
                },
                {
                        name:     "Reverse sorted",
                        input:    []int{5, 4, 3, 2, 1},
                        expected: []int{1, 2, 3, 4, 5},
                },
                {
                        name:     "Random order",
                        input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
                        expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
                },
                {
                        name:     "Duplicate elements",
                        input:    []int{3, 3, 2, 2, 1, 1},
                        expected: []int{1, 1, 2, 2, 3, 3},
                },
                {
                        name:     "Negative numbers",
                        input:    []int{-3, 0, 2, -5, 1, -4},
                        expected: []int{-5, -4, -3, 0, 1, 2},
                },
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        result := oddEvenSort(tt.input)
                        if !reflect.DeepEqual(result, tt.expected) {
                                t.Errorf("oddEvenSort(%v) = %v, want %v", tt.input, result, tt.expected)
                        }
                })
        }
}
