package shellsort

/*
Shell Sort is an in-place comparison sort algorithm that generalizes the insertion sort algorithm.
It starts by sorting pairs of elements far apart from each other, then progressively reduces the gap between elements to be compared.

See [Wikipedia](https://en.wikipedia.org/wiki/Shellsort)

Visual representation of the process:

    Initial array:  [9, 8, 3, 7, 5, 6, 4, 1]

    Gap = 4:        [9, 8, 3, 7, 5, 6, 4, 1]
                     |     |     |     |
                     v     v     v     v
                    [1, 8, 3, 7, 5, 6, 4, 9]

    Gap = 2:        [1, 8, 3, 7, 5, 6, 4, 9]
                     |  |  |  |  |  |  |  |
                     v  v  v  v  v  v  v  v
                    [1, 3, 4, 6, 5, 7, 8, 9]

    Gap = 1:        [1, 3, 4, 6, 5, 7, 8, 9]
                     | | | | | | | |
                     v v v v v v v v
    Final array:    [1, 3, 4, 5, 6, 7, 8, 9]
*/

func shellSort(arr []int) []int {
        n := len(arr)
        // Start with a large gap, then reduce the gap
        for gap := n / 2; gap > 0; gap /= 2 {
                // Do a gapped insertion sort for this gap size
                for i := gap; i < n; i++ {
                        // Add arr[i] to the elements that have been gap sorted
                        // Save arr[i] in temp and make a hole at position i
                        temp := arr[i]

                        // Shift earlier gap-sorted elements up until the correct location for arr[i] is found
                        var j int
                        for j = i; j >= gap && arr[j-gap] > temp; j -= gap {
                                arr[j] = arr[j-gap]
                        }

                        // Put temp (the original arr[i]) in its correct location
                        arr[j] = temp
                }
        }
        return arr
}

/*
Time Complexity:
- Best Case: O(n log n)
- Average Case: O(n^(4/3)) to O(n^(3/2))
- Worst Case: O(n^2)

Space Complexity: O(1) (in-place algorithm)

Note: The time complexity of Shell Sort depends on the gap sequence used.
The implementation above uses a simple gap sequence of n/2, n/4, ..., 1.
*/
