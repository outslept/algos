package bubblesort

// bubbleSort implements the Bubble Sort algorithm, which repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order. This process is repeated until the list is sorted.
func bubbleSort(arr []int) []int {
        n := len(arr)
        // Outer loop to make n-1 passes through the array.
        for i := 0; i < n-1; i++ {
                // Inner loop to compare adjacent elements and swap if needed.
                // The umber of comparisons reduces after each pass, as the largest element
                // "bubbles" to its correct position at the end of the array
                for j := 0; j < n-i-1; j++ {
                        // Swap if the current element is > the next element
                        if arr[j] > arr[j+1] {
                                arr[j], arr[j+1] = arr[j+1], arr[j]
                        }
                }
        }
        return arr
}
