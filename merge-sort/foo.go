package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mergeSort(arr []int) []int {
        if len(arr) <= 1 {
                return arr
        }

        mid := len(arr) / 2
        left := mergeSort(arr[:mid])
        right := mergeSort(arr[mid:])

        return merge(left, right)
}

func merge(left, right []int) []int {
        result := make([]int, 0, len(left)+len(right))
        i, j := 0, 0

        for i < len(left) && j < len(right) {
                if left[i] <= right[j] {
                        result = append(result, left[i])
                        i++
                } else {
                        result = append(result, right[j])
                        j++
                }
        }

        result = append(result, left[i:]...)
        result = append(result, right[j:]...)

        return result
}

func generateRandomArray(size int, maxValue int) []int {
        arr := make([]int, size)
        for i := 0; i < size; i++ {
                arr[i] = rand.Intn(maxValue)
        }
        return arr
}

func main() {
        rand.Seed(time.Now().UnixNano())

        size := 100000
        maxValue := 1000000

        arr := generateRandomArray(size, maxValue)

        fmt.Printf("Arr size: %d\n", size)
        fmt.Printf("Max: %d\n", maxValue)
        fmt.Printf("First 10 elements of input arr: %v\n", arr[:10])

        start := time.Now()
        sortedArr := mergeSort(arr)
        duration := time.Since(start)

        fmt.Printf("First 10 elements of sorted array: %v\n", sortedArr[:10])
        fmt.Printf("Exec time: %v\n", duration)

        fmt.Printf("Performance: %.2f elements/second\n", float64(size)/duration.Seconds())
        fmt.Printf("Avg time for elements: %.2f nanosec\n", float64(duration.Nanoseconds())/float64(size))
}
