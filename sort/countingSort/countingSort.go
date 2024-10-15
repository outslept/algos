package countingsort

/*
Counting Sort is a non-comparison-based sorting algorithm that works well when
there is limited range of input values. It operates by counting the number of
objects having distinct key values, and using arithmetic to determine the positions
of each key value in the output sequence.

See [Wikipedia](https://en.wikipedia.org/wiki/Counting_sort)

Visual representation of the process:

    Input array:   [4, 2, 2, 8, 3, 3, 1]

    Count array:   [0, 1, 2, 2, 1, 0, 0, 0, 1]
                    0  1  2  3  4  5  6  7  8

    Cumulative:    [0, 1, 3, 5, 6, 6, 6, 6, 7]
                    0  1  2  3  4  5  6  7  8

    Output array:  [1, 2, 2, 3, 3, 4, 8]
*/

func countingSort(arr []int) []int {
        if len(arr) == 0 {
                return arr
        }

        // Find the maximum element in the input array
        max := arr[0]
        for _, num := range arr {
                if num > max {
                        max = num
                }
        }

        // Create a count array to store the count of each unique object
        count := make([]int, max+1)
        for _, num := range arr {
                count[num]++
        }

        // Modify the count array to store actual position of elements in output array
        for i := 1; i <= max; i++ {
                count[i] += count[i-1]
        }

        // Build the output array
        output := make([]int, len(arr))
        for i := len(arr) - 1; i >= 0; i-- {
                output[count[arr[i]]-1] = arr[i]
                count[arr[i]]--
        }

        return output
}

/*
Time Complexity:
- Best Case: O(n + k), where n is the number of elements and k is the range of input
- Average Case: O(n + k)
- Worst Case: O(n + k)

Space Complexity: O(n + k)
*/
