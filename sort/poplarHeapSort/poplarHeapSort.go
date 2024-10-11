package poplarheapsort

func poplarHeapSort(arr []int) []int {
	n := len(arr)
	buildPoplarHeap(arr)
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		siftDown(arr, 0, i)
	}
	return arr
}

func buildPoplarHeap(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		siftDown(arr, i, len(arr))
	}
}

func siftDown(arr []int, root, n int) {
	for {
		child := 2*root + 1
		if child >= n {
			break
		}
		if child+1 < n && arr[child] < arr[child+1] {
			child++
		}
		if arr[root] >= arr[child] {
			break
		}
		arr[root], arr[child] = arr[child], arr[root]
		root = child
	}
}
