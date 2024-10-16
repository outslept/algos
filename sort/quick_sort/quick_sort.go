package quicksort

import (
        "math/rand"
)


/*
Quicksort is a sorting algorithm that uses a divide-and-conquer strategy.
It selects a 'pivot' element and partitions the array into two sub-arrays:
elements less than the pivot and elements greater than the pivot.
The algorithm then recursively sorts these sub-arrays.

See [Wikipedia](https://en.wikipedia.org/wiki/Quicksort)

Visual representation of the process:

    (4)---(2)---(7)
     |   /   \   |
     |  /     \  |
     | /       \ |
     |/         \|
    (1)---(5)---(3)

    1. Choose pivot (e.g., 4)
    2. Partition: [2, 1, 3] [4] [7, 5]
    3. Recursively sort sub-arrays
    4. Combine: [1, 2, 3, 4, 5, 7]
*/


func quickSort(arr []int) []int {
        // Base case: if the array has 1 or fewer elements, it's already sorted
        if len(arr) <= 1 {
                return arr
        }

        // Choose a random pivot to improve average-case performance
        pivotIndex := rand.Intn(len(arr))
        pivot := arr[pivotIndex]

        // Create three sub-arrays:
        // - left:  elements less than pivot
        // - equal: elements equal to pivot (handles duplicates)
        // - right: elements greater than pivot
        var left, equal, right []int

        // Partition the array
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


        /*
        Visualization of partitioning:

            Before:  [4, 2, 7, 1, 5, 3]  (pivot = 4)
                      ^
            After:   [2, 1, 3] [4] [7, 5]
                      left   equal right

                (2)---(1)   (4)   (7)---(5)
                 |     |     |     |     |
                 |     |     |     |     |
                 |     |     |     |     |
                (3)    |     |     |    (3)
        */


        // Recursively apply quickSort to left and right sub-arrays
        left = quickSort(left)
        right = quickSort(right)

        // Combine the sorted sub-arrays and pivot elements
        return append(append(left, equal...), right...)
}

/*
Note: This implementation is not in-place. An in-place version
would modify the original array without creating new slices.

Time Complexity:
- Average and Best Case: O(n log n)
- Worst Case: O(n^2) (rare with random pivot selection)

Space Complexity:
- Average Case: O(log n)
- Worst Case: O(n)
*/
