package mergesort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
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
                        name:     "Mixed order",
                        input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
                        expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
                },
                {
                        name:     "Duplicate elements",
                        input:    []int{3, 3, 3, 1, 1, 2, 2},
                        expected: []int{1, 1, 2, 2, 3, 3, 3},
                },
                {
                        name:     "Negative numbers",
                        input:    []int{-5, 2, -8, 10, -3, 0, 7},
                        expected: []int{-8, -5, -3, 0, 2, 7, 10},
                },
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        result := mergeSort(tt.input)
                        if !reflect.DeepEqual(result, tt.expected) {
                                t.Errorf("mergeSort() = %v, want %v", result, tt.expected)
                        }
                })
        }
}

func TestMerge(t *testing.T) {
        tests := []struct {
                name     string
                left     []int
                right    []int
                expected []int
        }{
                {
                        name:     "Both empty",
                        left:     []int{},
                        right:    []int{},
                        expected: []int{},
                },
                {
                        name:     "Left empty",
                        left:     []int{},
                        right:    []int{1, 2, 3},
                        expected: []int{1, 2, 3},
                },
                {
                        name:     "Right empty",
                        left:     []int{1, 2, 3},
                        right:    []int{},
                        expected: []int{1, 2, 3},
                },
                {
                        name:     "Equal length",
                        left:     []int{1, 3, 5},
                        right:    []int{2, 4, 6},
                        expected: []int{1, 2, 3, 4, 5, 6},
                },
                {
                        name:     "Left longer",
                        left:     []int{1, 3, 5, 7},
                        right:    []int{2, 4},
                        expected: []int{1, 2, 3, 4, 5, 7},
                },
                {
                        name:     "Right longer",
                        left:     []int{1, 3},
                        right:    []int{2, 4, 6, 8},
                        expected: []int{1, 2, 3, 4, 6, 8},
                },
                {
									name:     "With duplicates",
									left:     []int{1, 2, 2, 3},
									right:    []int{2, 3, 4, 4},
									expected: []int{1, 2, 2, 2, 3, 3, 4, 4},
					},
					{
									name:     "Negative numbers",
									left:     []int{-5, -3, 0, 2},
									right:    []int{-4, -2, 1, 3},
									expected: []int{-5, -4, -3, -2, 0, 1, 2, 3},
					},
	}

	for _, tt := range tests {
					t.Run(tt.name, func(t *testing.T) {
									result := merge(tt.left, tt.right)
									if !reflect.DeepEqual(result, tt.expected) {
													t.Errorf("merge() = %v, want %v", result, tt.expected)
									}
					})
	}
}

// Benchmark function for mergeSort
func BenchmarkMergeSort(b *testing.B) {
	input := []int{5, 2, 8, 12, 1, 6}
	for i := 0; i < b.N; i++ {
					mergeSort(input)
	}
}

// Benchmark function for merge
func BenchmarkMerge(b *testing.B) {
	left := []int{1, 3, 5}
	right := []int{2, 4, 6}
	for i := 0; i < b.N; i++ {
					merge(left, right)
	}
}
