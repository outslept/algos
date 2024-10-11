package quicksort

import (
	"math/rand"
)

// quickSort implement the QuickSort algorithm (https://en.wikipedia.org/wiki/Quicksort), which is a divide-and conquer sorting technique. It recursively partitions the input array based on a pivot element, and sorts the subarrays to the left (elements < the pivot) and right (elements > the pivot).
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
