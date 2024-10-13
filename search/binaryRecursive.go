package search

func BinarySearchRecursive(array []int, target int) int {
	return binarySearchRecursiveHelper(array, target, 0, len(array)-1)
}

func binarySearchRecursiveHelper(array []int, target, start, end int) int {
	if len(array) == 0 {
			return -1
	}

	if target < array[start] || target > array[end] {
			return -1
	}

	middle := (start + end) >> 1

	if array[middle] == target {
			return middle
	}

	if start > end {
			return -1
	}

	if target < array[middle] {
			return binarySearchRecursiveHelper(array, target, start, middle-1)
	}

	return binarySearchRecursiveHelper(array, target, middle+1, end)
}
