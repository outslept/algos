package estimatePi

import (
	"math/rand"
	"time"
)

func estimatePi(iterations int) float64 {
	rand.Seed(time.Now().UnixNano())
	inside := 0
	for i := 0; i < iterations; i++ {
			x := rand.Float64()
			y := rand.Float64()
			if x*x+y*y <= 1 {
					inside++
			}
	}
	return 4 * float64(inside) / float64(iterations)
}
