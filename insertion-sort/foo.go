package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(arr []int) []int {
        n := len(arr)
        for i := 0; i < n-1; i++ {
                for j := 0; j < n-i-1; j++ {
                        if arr[j] > arr[j+1] {
                                arr[j], arr[j+1] = arr[j+1], arr[j]
                        }
                }
        }
        return arr
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

        size := 10000
        maxValue := 1000000

        arr := generateRandomArray(size, maxValue)

        fmt.Printf("Arr size: %d\n", size)
        fmt.Printf("Max: %d\n", maxValue)
        fmt.Printf("First 10 elements of input arr: %v\n", arr[:10])

        start := time.Now()
        sortedArr := bubbleSort(arr)
        duration := time.Since(start)

        fmt.Printf("First 10 elements of sorted arr: %v\n", sortedArr[:10])
        fmt.Printf("Exec time: %v\n", duration)

        fmt.Printf("Performance: %.2f elements/second\n", float64(size)/duration.Seconds())
        fmt.Printf("Avg time for elements: %.2f nanosec\n", float64(duration.Nanoseconds())/float64(size))
}
