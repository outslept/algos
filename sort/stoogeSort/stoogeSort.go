package stoogesort


/*
Stooge Sort is a recursive sorting algorithm with a unique approach.
It divides the array into three parts and recursively sorts the first 2/3,
then the last 2/3, and finally the first 2/3 again.

See [Wikipedia](https://en.wikipedia.org/wiki/Stooge_sort)

Visual representation of the process:

    (5)---(2)---(8)---(1)---(9)---(3)
     |     |     |     |     |     |
     |     |     |     |     |     |
    (2)---(5)---(8)---(1)---(9)---(3)
     |           |           |
     |           |           |
    (2)---(5)---(1)---(8)---(9)---(3)
     |     |     |     |     |     |
     |     |     |     |     |     |
    (1)---(2)---(5)---(3)---(8)---(9)

    1. Sort first 2/3
    2. Sort last 2/3
    3. Sort first 2/3 again
    4. Repeat recursively
*/


func stoogeSort(arr []int, i, j int) []int {
	// If the first element is larger than the last, swap them
	if arr[i] > arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
	}

	// If there are more than 2 elements in the array
	if j-i+1 > 2 {
					t := (j - i + 1) / 3 // Calculate the size of 1/3 of the current subarray

					// Recursively sort the first 2/3 of the array
					stoogeSort(arr, i, j-t)

					// Recursively sort the last 2/3 of the array
					stoogeSort(arr, i+t, j)

					// Recursively sort the first 2/3 again to ensure all elements are in order
					stoogeSort(arr, i, j-t)
	}
	return arr
}

/*
Time Complexity:
- Worst, Average, and Best Case: O(n^(log 3 / log 1.5)) ≈ O(n^2.7095)

Space Complexity:
- O(log n) due to the recursion stack

Note: Stooge Sort is primarily of theoretical interest and is not used in practice
due to its poor time complexity. It's less efficient than even simple algorithms.
*/
