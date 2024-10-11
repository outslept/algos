package stoogesort

func stoogeSort(arr []int, i, j int) []int {
	if arr[i] > arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
	}
	if j-i+1 > 2 {
					t := (j - i + 1) / 3
					stoogeSort(arr, i, j-t)
					stoogeSort(arr, i+t, j)
					stoogeSort(arr, i, j-t)
	}
	return arr
}
