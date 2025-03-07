package sort

import (
	"math/rand"
	"time"
)

func BogoSort(arr []int) []int {
	rand.Seed(time.Now().UnixNano())
	for !isSorted(arr) {
		shuffle(arr)
	}
}

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func shuffle(arr []int) {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
