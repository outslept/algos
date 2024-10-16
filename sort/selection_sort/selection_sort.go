package selectionsort

/*
Selection Sort is a simple comparison-based sorting algorithm.
It divides the input list into two parts: a sorted portion at the left end
and an unsorted portion at the right end. Initially, the sorted portion is empty
and the unsorted portion is the entire list.

The algorithm repeatedly selects the smallest (or largest) element from the
unsorted portion and moves it to the end of the sorted portion.

See [Wikipedia](https://en.wikipedia.org/wiki/Selection_sort)

Visual representation of the process:

    Initial array:  [64, 25, 12, 22, 11]

    First pass:     [11, 25, 12, 22, 64]
                     ^   |   |   |   |
                     |   |   |   |   |
                     +---+---+---+---+

    Second pass:    [11, 12, 25, 22, 64]
                         ^   |   |   |
                         |   |   |   |
                         +---+---+---+

    Third pass:     [11, 12, 22, 25, 64]
                             ^   |   |
                             |   |   |
                             +---+---+

    Fourth pass:    [11, 12, 22, 25, 64]
                                 ^   |
                                 |   |
                                 +---+

    Final array:    [11, 12, 22, 25, 64]
*/

func selectionSort(arr []int) []int {
        n := len(arr)
        for i := 0; i < n; i++ {
                // Assume the current index is the minimum
                minIdx := i

                // Find the minimum element in the unsorted portion
                for j := i + 1; j < n; j++ {
                        if arr[j] < arr[minIdx] {
                                minIdx = j
                        }
                }

                // Swap the found minimum element with the first element of the unsorted portion
                arr[i], arr[minIdx] = arr[minIdx], arr[i]
        }
        return arr
}

/*
Time Complexity:
- Best Case: O(n^2)
- Average Case: O(n^2)
- Worst Case: O(n^2)

Space Complexity: O(1) (in-place algorithm)
*/
