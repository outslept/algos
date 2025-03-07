package sort

func QuickSort(arr []int) []int {
        if len(arr) <= 1 {
                return arr
        }

        quickSortHelper(arr, 0, len(arr)-1)
        return arr
}

func quickSortHelper(arr []int, low, high int) {
        if low < high {
                pivot := partition(arr, low, high)
                quickSortHelper(arr, low, pivot-1)
                quickSortHelper(arr, pivot+1, high)
        }
}

func partition(arr []int, low, high int) int {
        pivot := arr[high]
        i := low - 1

        for j := low; j < high; j++ {
                if arr[j] <= pivot {
                        i++
                        arr[i], arr[j] = arr[j], arr[i]
                }
        }

        arr[i+1], arr[high] = arr[high], arr[i+1]
        return i + 1
}
