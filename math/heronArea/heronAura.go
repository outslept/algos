package heronarea

import "math"

func heronArea(a, b, c float64) float64 {
    s := (a + b + c) / 2
    return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}
