package mergesort

// mergeSort implements the Merge Sort algorithm, a divide-and-conquer sorting technique. It recursively splits the input array into smaller subarrays back together in sorted order.
func mergeSort(arr []int) []int {
        // Base case:if the array has 1 or fewer elements, it is already sorted.
        if len(arr) <= 1 {
                return arr
        }

        // Recursively split the array into two halves.
        mid := len(arr) / 2
        left := mergeSort(arr[:mid]) // Sort the left half
        right := mergeSort(arr[mid:]) // Sort the right half

        // Merge the sorted halves back together.
        return merge(left, right)
}

// merge takes two sorted array(`left` and `right`) and merges them into a single sorted array.
func merge(left, right []int) []int {
        // Preallocate a result slice with a capacity equal to the combined size of the `left` and `right`.
        result := make([]int, 0, len(left)+len(right))
        
        // Initialize two indices, `i` for `left` and `j` for `right`
        i, j := 0, 0

        // Merge the two arrays by comparing the elements at each index and appending the smaller one.
        for i < len(left) && j < len(right) {
                if left[i] <= right[j] {
                        result = append(result, left[i])
                        i++
                } else {
                        result = append(result, right[j])
                        j++
                }
        }

        // Append any remaining elements from `left` (if any)
        result = append(result, left[i:]...)
        
        // Append any remaining elements from `right` (if any)
        result = append(result, right[j:]...)

        return result
}
