package timsort

import (
    "math"
)

/*
TimSort is a hybrid sorting algorithm derived from merge sort and insertion sort.
It was designed to perform well on many kinds of real-world data.

The algorithm divides the array into small runs, sorts these runs using insertion sort,
and then merges the runs using a merge sort approach.

See [Wikipedia](https://en.wikipedia.org/wiki/Timsort)

Visual representation of the process:

    [7, 3, 4, 1, 5, 2, 6, 8]
           /           \
    [3, 4, 7]    [1, 2, 5, 6, 8]
      /    \        /        \
   [3, 4] [7]    [1, 2]    [5, 6, 8]
     |      |      |          |
   Insert  Insert Insert    Insert
     Sort   Sort   Sort      Sort
      |      |      |          |
    Merge   Merge  Merge     Merge
      \    /        \        /
    [3, 4, 7]    [1, 2, 5, 6, 8]
           \           /
        Final Merge
             |
    [1, 2, 3, 4, 5, 6, 7, 8]
*/

const (
    minRun = 32 // Minimum size of a run
)

func TimSort(arr []int) {
    n := len(arr)

    // Sort small arrays using insertion sort
    if n < minRun {
        insertionSort(arr, 0, n)
        return
    }

    // Divide array into runs and sort them
    for i := 0; i < n; i += minRun {
        end := int(math.Min(float64(i+minRun), float64(n)))
        insertionSort(arr, i, end)
    }

    // Merge runs
    for size := minRun; size < n; size *= 2 {
        for start := 0; start < n; start += size * 2 {
            mid := start + size
            end := int(math.Min(float64(start+size*2), float64(n)))
            if mid < end {
                merge(arr, start, mid, end)
            }
        }
    }
}

/*
Insertion Sort is used for sorting small runs.
It's efficient for small arrays and partially sorted arrays.

Visual representation:

    [5, 2, 4, 6, 1, 3]
     ^  ^
    [2, 5, 4, 6, 1, 3]
        ^  ^
    [2, 4, 5, 6, 1, 3]
           ^  ^
    [2, 4, 5, 6, 1, 3]
              ^  ^
    [1, 2, 4, 5, 6, 3]
                 ^  ^
    [1, 2, 3, 4, 5, 6]
*/
func insertionSort(arr []int, left, right int) {
    for i := left + 1; i < right; i++ {
        key := arr[i]
        j := i - 1
        for j >= left && arr[j] > key {
            arr[j+1] = arr[j]
            j--
        }
        arr[j+1] = key
    }
}

/*
Merge function combines two sorted subarrays into one sorted array.

Visual representation:

    [3, 4, 7] [1, 2, 5]
     ^         ^
    [1, 3, 4, 7] [2, 5]
         ^        ^
    [1, 2, 3, 4, 7] [5]
            ^         ^
    [1, 2, 3, 4, 5, 7]
*/
func merge(arr []int, left, mid, right int) {
	leftArr := make([]int, mid-left)
	rightArr := make([]int, right-mid)

	// Copy data to temporary arrays
	copy(leftArr, arr[left:mid])
	copy(rightArr, arr[mid:right])

	i, j, k := 0, 0, left

	// Merge temporary arrays back into arr[left..right]
	for i < len(leftArr) && j < len(rightArr) {
			if leftArr[i] <= rightArr[j] {
					arr[k] = leftArr[i]
					i++
			} else {
					arr[k] = rightArr[j]
					j++
			}
			k++
	}

	// Copy remaining elements of leftArr[], if any
	for i < len(leftArr) {
			arr[k] = leftArr[i]
			i++
			k++
	}

	// Copy remaining elements of rightArr[], if any
	for j < len(rightArr) {
			arr[k] = rightArr[j]
			j++
			k++
	}
}

/*
Tim Sort combines the strengths of Merge Sort and Insertion Sort:
- Uses Insertion Sort for small runs (efficient for small, partially sorted arrays)
- Uses Merge Sort strategy for combining runs (efficient for larger datasets)

Time Complexity:
- Best Case: O(n) when the array is already sorted
- Average and Worst Case: O(n log n)

Space Complexity:
- O(n) due to the temporary arrays used in merging
*/

