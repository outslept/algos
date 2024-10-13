package search

func BinarySearchIterative(array []int, target int) int {
	if len(array) == 0 {
		return -1
	}

	start := 0
	end := len(array) - 1

	if target < array[start] || target > array[end] {
		return -1
	}

	for start <= end {
		middle := (start + end) >> 1

		if array[middle] == target {
			return middle
		}

		if target < array[middle] {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	
	return -1
}
