package cyclesort

func cycleSort(arr []int) []int {
	n := len(arr)
	for cycleStart := 0; cycleStart < n-1; cycleStart++ {
					item := arr[cycleStart]
					pos := cycleStart
					for i := cycleStart + 1; i < n; i++ {
									if arr[i] < item {
													pos++
									}
					}
					if pos == cycleStart {
									continue
					}
					for item == arr[pos] {
									pos++
					}
					arr[pos], item = item, arr[pos]
					for pos != cycleStart {
									pos = cycleStart
									for i := cycleStart + 1; i < n; i++ {
													if arr[i] < item {
																	pos++
													}
									}
									for item == arr[pos] {
													pos++
									}
									arr[pos], item = item, arr[pos]
					}
	}
	return arr
}
