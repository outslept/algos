package twopointers

func TwoPointers(nums []int, target int) [2]int {
	left, right := 0, len(nums)-1

	for left < right {
		currentSum := nums[left] + nums[right]

		if currentSum == target {
			return [2]int{left, right}
		}

		if currentSum < target {
			left++
		} else {
			right--
		}
	}

	return [2]int{-1, -1}
}
