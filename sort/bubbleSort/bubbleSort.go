package bubblesort

/*
Bubble Sort is a simple sorting algorithm that repeatedly steps through the list,
compares adjacent elements and swaps them if they are in the wrong order.
The pass through the list is repeated until the list is sorted.

See [Wikipedia](https://en.wikipedia.org/wiki/Bubble_sort)

Visual representation of the process:

    (5)---(2)---(8)---(1)---(9)
     |     |     |     |     |
     v     v     v     v     v
    (2)---(5)---(1)---(8)---(9)
     |     |     |     |     |
     v     v     v     v     v
    (2)---(1)---(5)---(8)---(9)
     |     |     |     |     |
     v     v     v     v     v
    (1)---(2)---(5)---(8)---(9)

    1. Compare adjacent elements
    2. Swap if they are in the wrong order
    3. Repeat until no more swaps are needed
*/

func bubbleSort(arr []int) []int {
    n := len(arr)
    
    // Outer loop to make n-1 passes through the array
    for i := 0; i < n-1; i++ {
        // Inner loop to compare adjacent elements and swap if needed
        // The number of comparisons reduces after each pass, as the largest element
        // "bubbles" to its correct position at the end of the array
        for j := 0; j < n-i-1; j++ {
            // Swap if the current element is greater than the next element
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }

    return arr
}

/*
Time Complexity:
- Worst and Average Case: O(n^2), where n is the number of elements
- Best Case: O(n) when the array is already sorted

Space Complexity: O(1) as it sorts in-place
*/
