package oddevensort


/*
Odd-Even Sort is a relatively simple sorting algorithm.
It works by repeatedly making two types of comparisons:
1. Odd-indexed elements with their next even-indexed neighbor
2. Even-indexed elements with their next odd-indexed neighbor

This process continues until no more swaps are needed, indicating the array is sorted.

See [Wikipedia](https://en.wikipedia.org/wiki/Odd%E2%80%93even_sort)

Visual representation of the process:

    (5)---(2)---(8)---(1)---(9)---(3)
     |     |     |     |     |     |
     v     v     v     v     v     v
    (2)---(5)---(1)---(8)---(3)---(9)
     |     |     |     |     |     |
     v     v     v     v     v     v
    (2)---(1)---(5)---(3)---(8)---(9)
     |     |     |     |     |     |
     v     v     v     v     v     v
    (1)---(2)---(3)---(5)---(8)---(9)

    1. Compare and swap odd-indexed elements with their next even-indexed neighbor
    2. Compare and swap even-indexed elements with their next odd-indexed neighbor
    3. Repeat until no more swaps are needed
*/


func oddEvenSort(arr []int) []int {
	sorted := false
	n := len(arr)

	for !sorted {
					sorted = true

					// Odd-indexed comparisons
					for i := 1; i < n-1; i += 2 {
									if arr[i] > arr[i+1] {
													arr[i], arr[i+1] = arr[i+1], arr[i]
													sorted = false
									}
					}

					// Even-indexed comparisons
					for i := 0; i < n-1; i += 2 {
									if arr[i] > arr[i+1] {
													arr[i], arr[i+1] = arr[i+1], arr[i]
													sorted = false
									}
					}
	}
	return arr
}

/*
Time Complexity:
- Worst and Average Case: O(n^2)
- Best Case: O(n) when the array is already sorted

Space Complexity: O(1) as it sorts in-place

Note: This algorithm is not efficient for large datasets but can be useful
in parallel computing environments due to its locality of comparison.
*/
