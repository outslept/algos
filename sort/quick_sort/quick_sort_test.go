package quicksort

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
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
                        rand.Seed(42)
                        result := quickSort(tt.input)
                        if !reflect.DeepEqual(result, tt.expected) {
                                t.Errorf("quickSort() = %v, want %v", result, tt.expected)
                        }
                })
        }
}

func TestQuickSortLarge(t *testing.T) {
        rand.Seed(42)
        size := 10000
        input := make([]int, size)
        for i := 0; i < size; i++ {
                input[i] = rand.Intn(10000) - 5000 // Random numbers between -5000 and 4999
        }

        result := quickSort(input)

        for i := 1; i < len(result); i++ {
                if result[i] < result[i-1] {
                        t.Errorf("quickSort failed to sort large input: element at index %d is smaller than the previous element", i)
                        return
                }
        }
}

func TestQuickSortStability(t *testing.T) {
        input := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
        expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}

        for i := 0; i < 100; i++ {
                rand.Seed(int64(i)) 
                result := quickSort(input)
                if !reflect.DeepEqual(result, expected) {
                        t.Errorf("quickSort() on run %d = %v, want %v", i, result, expected)
                        return
                }
        }
			}

			func BenchmarkQuickSort(b *testing.B) {
							rand.Seed(42)
							size := 1000
							input := make([]int, size)
							for i := 0; i < size; i++ {
											input[i] = rand.Intn(1000)
							}
			
							b.ResetTimer()
							for i := 0; i < b.N; i++ {
											quickSort(input)
							}
			}
			
			func TestQuickSortWithAllEqualElements(t *testing.T) {
							input := []int{5, 5, 5, 5, 5, 5, 5}
							expected := []int{5, 5, 5, 5, 5, 5, 5}
							
							result := quickSort(input)
							if !reflect.DeepEqual(result, expected) {
											t.Errorf("quickSort() with all equal elements = %v, want %v", result, expected)
							}
			}
			
			func TestQuickSortWithRepeatedCalls(t *testing.T) {
							input := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
							expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
			
							for i := 0; i < 5; i++ {
											result := quickSort(input)
											if !reflect.DeepEqual(result, expected) {
															t.Errorf("quickSort() on call %d = %v, want %v", i, result, expected)
											}
							}
			}
			
			func TestQuickSortWithEdgeCases(t *testing.T) {
							tests := []struct {
											name     string
											input    []int
											expected []int
							}{
											{
															name:     "Slice with two elements",
															input:    []int{2, 1},
															expected: []int{1, 2},
											},
											{
															name:     "Slice with repeated elements at the start",
															input:    []int{1, 1, 1, 2, 3, 4, 5},
															expected: []int{1, 1, 1, 2, 3, 4, 5},
											},
											{
															name:     "Slice with repeated elements at the end",
															input:    []int{1, 2, 3, 4, 5, 5, 5},
															expected: []int{1, 2, 3, 4, 5, 5, 5},
											},
											{
															name:     "Slice with alternating elements",
															input:    []int{1, 2, 1, 2, 1, 2},
															expected: []int{1, 1, 1, 2, 2, 2},
											},
							}
			
							for _, tt := range tests {
											t.Run(tt.name, func(t *testing.T) {
															rand.Seed(42) // Set a fixed seed for reproducibility
															result := quickSort(tt.input)
															if !reflect.DeepEqual(result, tt.expected) {
																			t.Errorf("quickSort() = %v, want %v", result, tt.expected)
															}
											})
							}
			}
