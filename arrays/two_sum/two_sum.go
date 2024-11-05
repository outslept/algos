package twosum

func TwoSum(nums []int, target int) [2]int {
	numMap := make(map[int]int, len(nums))

	for i, num := range nums {
		if j, exists := numMap[target-num]; exists {
			return [2]int{j, i}
		}

		numMap[num] = i
	}

	return [2]int{-1, -1}
}
