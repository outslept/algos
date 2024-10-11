package countingsort


func countingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	count := make([]int, max+1)
	for _, num := range arr {
		count[num]++
	}

	for i:= 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	output := make([]int, len(arr))
	for i:= len(arr) - 1; i >= 0; i-- {
		output[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}

	return output
}
