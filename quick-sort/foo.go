package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(arr []int) []int {
        if len(arr) <= 1 {
                return arr
        }

        pivot := arr[rand.Intn(len(arr))]
        var left, right, equal []int

        for _, num := range arr {
                switch {
                case num < pivot:
                        left = append(left, num)
                case num == pivot:
                        equal = append(equal, num)
                case num > pivot:
                        right = append(right, num)
                }
        }

        left = quickSort(left)
        right = quickSort(right)

        return append(append(left, equal...), right...)
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
        sortedArr := quickSort(arr)
        duration := time.Since(start)

        fmt.Printf("First 10 elements of sorted arr: %v\n", sortedArr[:10])
        fmt.Printf("Exec time: %v\n", duration)

        fmt.Printf("Performance: %.2f elements/second\n", float64(size)/duration.Seconds())
        fmt.Printf("Avg time for elements: %.2f nanosec\n", float64(duration.Nanoseconds())/float64(size))
}
