package gnomesort

func gnomeSort(arr []int) []int {
	i := 1
	for i < len(arr) {
		if i == 0 || arr[i-1] <= arr[i] {
			i++
		} else {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			i--
		}
	}
	return arr
}
