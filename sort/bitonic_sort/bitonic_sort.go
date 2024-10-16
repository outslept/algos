package bitonicsort

/*
Bitonic Sort is a comparison-based sorting algorithm that can be run in parallel.
It focuses on converting a random sequence of numbers into a bitonic sequence,
then sorting that sequence into a monotonically increasing sequence.

A sequence is called bitonic if it first increases, then decreases.

See [Wikipedia](https://en.wikipedia.org/wiki/Bitonic_sorter)

Visual representation of the process:

    (3)---(7)---(4)---(8)---(6)---(2)---(1)---(5)
     |     |     |     |     |     |     |     |
     v     v     v     v     v     v     v     v
    (3)---(7)---(4)---(8)---(1)---(2)---(6)---(5)
     |     |     |     |     |     |     |     |
     v     v     v     v     v     v     v     v
    (3)---(4)---(7)---(8)---(1)---(2)---(5)---(6)
     |     |     |     |     |     |     |     |
     v     v     v     v     v     v     v     v
    (1)---(2)---(3)---(4)---(5)---(6)---(7)---(8)

    1. Create bitonic sequences
    2. Merge bitonic sequences
    3. Repeat until the entire sequence is sorted
*/

func bitonicSort(arr []int) []int {
    n := len(arr)
    k := 2

    // Outer loop: double the size of subarrays in each iteration
    for k <= n {
        // Middle loop: halve the size of comparison in each iteration
        for j := k / 2; j > 0; j /= 2 {
            // Inner loop: compare elements
            for i := 0; i < n; i++ {
                l := i ^ j  // XOR operation to find the element to compare with

                // Compare and swap if necessary
                if l > i {
                    if (i&k) == 0 && arr[i] > arr[l] || (i&k) != 0 && arr[i] < arr[l] {
                        arr[i], arr[l] = arr[l], arr[i]
                    }
                }
            }
        }
        k *= 2
    }

    return arr
}

/*
Time Complexity:
- All cases: O(n log^2 n), where n is the number of elements

Space Complexity: O(1) as it sorts in-place

Note: It works most efficiently when n is a power of 2. For other sizes, the array can be
padded with large values to reach the next power of 2.
*/
