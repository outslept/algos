package bitonicsort

func bitonicSort(arr []int) []int {
	n := len(arr)
	k := 2
	for k <= n {
					for j := k / 2; j > 0; j /= 2 {
									for i := 0; i < n; i++ {
													l := i ^ j
													if l > i {
																	if (i&k) == 0 && arr[i] > arr[l] || (i&k) != 0 && arr[i] < arr[l] {
																					arr[i], arr[l] = arr[l], arr[i]
																	}
													}
									}
					}
					k *= 2
	}
	return arr
}
