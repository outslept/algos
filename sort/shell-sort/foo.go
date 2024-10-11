package shellsort

func shellSort(arr []int) []int {
	n := len(arr)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
	}
	return arr
}
