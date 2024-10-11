package distance

import "math"

func distance(x1, y1, x2, y2 float64) float64 {
    dx := x2 - x1
    dy := y2 - y1
    return math.Sqrt(dx*dx + dy*dy)
}
