package combsort

/*
Comb Sort is an improvement over Bubble Sort. It eliminates small values
near the end of the list which are known to slow down Bubble Sort (the so-called "turtle" values).
The basic idea is to eliminate turtles or small values at the end of the list,
since in a bubble sort these slow the sorting down tremendously.

See [Wikipedia](https://en.wikipedia.org/wiki/Comb_sort)

Visual representation of the process:

    Initial array:  [8, 4, 1, 56, 3, -44, 23, -6, 28, 0]

    Gap = 10/1.3 ≈ 7:  [8,  4,  1, 56,  3, -44, 23, -6, 28,  0]
                        |                   |
                        +-------------------+
                       [0,  4,  1, 56,  3, -44, 23, -6, 28,  8]

    Gap = 7/1.3 ≈ 5:   [0,  4,  1, 56,  3, -44, 23, -6, 28,  8]
                        |              |
                        +--------------+
                       [0, -44,  1, 56,  3,  4, 23, -6, 28,  8]

    Gap = 5/1.3 ≈ 3:   [0, -44,  1, 56,  3,  4, 23, -6, 28,  8]
                        |        |
                        +--------+
                       [-44, -6,  1, 56,  3,  4, 23,  0, 28,  8]

    ... continues until gap = 1 and no more swaps are needed ...

    Final array:    [-44, -6, 0, 1, 3, 4, 8, 23, 28, 56]
*/

func combSort(arr []int) []int {
        n := len(arr)
        gap := n
        shrink := 1.3
        sorted := false

        for !sorted {
                // Update the gap value
                gap = int(float64(gap) / shrink)
                if gap <= 1 {
                        gap = 1
                        sorted = true // Will be set to false if any swaps occur
                }

                // Compare elements with the current gap
                for i := 0; i+gap < n; i++ {
                        if arr[i] > arr[i+gap] {
                                // Swap elements
                                arr[i], arr[i+gap] = arr[i+gap], arr[i]
                                sorted = false
                        }
                }
        }
        return arr
}

/*
Time Complexity:
- Best Case: O(n log n)
- Average Case: O(n^2 / 2^p), where p is the number of increments
- Worst Case: O(n^2)

Space Complexity: O(1) (in-place algorithm)

Note: The shrink factor of 1.3 is found to be the most effective empirically (according to Wiki),
but other values can be used. Comb Sort can be seen as a generalization of
Bubble Sort (when gap=1) and Shell Sort (which uses a different gap sequence).
*/
