package cyclesort

/*
Cycle Sort is an in-place, unstable sorting algorithm that is particularly useful when memory writes are costly.
It minimizes the number of memory writes to sort an array. It's based on the idea that the permutation to be sorted
can be decomposed into cycles, and each cycle can be rotated to make a sorted array.

See [Wikipedia](https://en.wikipedia.org/wiki/Cycle_sort)

Visual representation of the process:

    Initial array:  [5, 2, 1, 4, 3]

    First cycle:    [1, 2, 5, 4, 3]
                     ^     ^
                     |     |
                     +-----+

    Second cycle:   [1, 2, 3, 4, 5]
                           ^     ^
                           |     |
                           +-----+

    Final array:    [1, 2, 3, 4, 5]
*/

func cycleSort(arr []int) []int {
        n := len(arr)
        // Loop through the array to find cycles to rotate
        for cycleStart := 0; cycleStart < n-1; cycleStart++ {
                item := arr[cycleStart]

                // Find where to put the item
                pos := cycleStart
                for i := cycleStart + 1; i < n; i++ {
                        if arr[i] < item {
                                pos++
                        }
                }

                // If the item is already in the correct position
                if pos == cycleStart {
                        continue
                }

                // Ignore all duplicate elements
                for item == arr[pos] {
                        pos++
                }

                // Put the item to its right position
                arr[pos], item = item, arr[pos]

                // Rotate rest of the cycle
                for pos != cycleStart {
                        pos = cycleStart
                        // Find where to put the item
                        for i := cycleStart + 1; i < n; i++ {
                                if arr[i] < item {
                                        pos++
                                }
                        }

                        // Ignore all duplicate elements
                        for item == arr[pos] {
                                pos++
                        }

                        // Put the item to its right position
                        arr[pos], item = item, arr[pos]
                }
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
