package geometricmean

import "math"

func geometricMean(numbers []float64) float64 {
    product := 1.0
    for _, num := range numbers {
        product *= num
    }
    return math.Pow(product, 1/float64(len(numbers)))
}
