// quicksort.go
// description: Implementation of in-place quicksort algorithm.
// details:
// A simple in-place quicksort algorithm implementation. [Wikipedia](https://en.wikipedia.org/wiki/Quicksort)
// see quick_sort_test.go for a test implementation.

package quicksort

import (
	"math/rand"
)

func quickSort(arr []int) []int {
        // Base case: if the array has 1 or fewer elements, it is already sorted.
        if len(arr) <= 1 {
                return arr
        }

        // Randomly select a pivot element from the array.
        pivot := arr[rand.Intn(len(arr))]

        // Partition the array into three subarrays:
        // - `left`: elements less than the pivot,
        // - `equal`: elements equal to the pivot (to handle duplicates),
        // - `right`: elements greater than the pivot.
        var left, right, equal []int

        for _, num := range arr {
                switch {
                case num < pivot:
                        left = append(left, num)
                case num == pivot:
                        equal = append(equal, num)
                case num > pivot:
                        right = append(right, num)
                }
        }

        // Recursively apply quickSort to the `left` and `right` subarrays.
        left = quickSort(left)
        right = quickSort(right)

        // Concat the sorted left, the equal pivot values, and the sorted right.
        return append(append(left, equal...), right...)
}
