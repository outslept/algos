package slidingwindow

func MaxSumSlidingWindow(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}

	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}

	maxSum := windowSum

	for i := k; i < len(nums); i++ {
		windowSum = windowSum - nums[i-k] + nums[i]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}
