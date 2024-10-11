
// oddevensort.go
// description: Implementation of odd-even sort algorithm.
// details:
// A simple sorting odd-even sort algorithm implementation. [Wikipedia](https://en.wikipedia.org/wiki/Odd%E2%80%93even_sort)
// see oddevensort_test.go for a test implementation.

package oddevensort

func oddEvenSort(arr []int) []int {
	sorted := false
	n := len(arr)
	for !sorted {
					sorted = true
					for i := 1; i < n-1; i += 2 {
									if arr[i] > arr[i+1] {
													arr[i], arr[i+1] = arr[i+1], arr[i]
													sorted = false
									}
					}
					for i := 0; i < n-1; i += 2 {
									if arr[i] > arr[i+1] {
													arr[i], arr[i+1] = arr[i+1], arr[i]
													sorted = false
									}
					}
	}
	return arr
}
