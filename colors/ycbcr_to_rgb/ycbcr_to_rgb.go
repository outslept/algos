package ycbcr_to_rgb

import "math"

func YCbCrToRGB(y, cb, cr uint8) (r, g, b uint8) {
	y1 := float64(y)
	cb1 := float64(cb)
	cr1 := float64(cr)

	r = uint8(clamp(y1 + 1.402*cr1))
	g = uint8(clamp(y1 + 0.344136*cb1 - 0.714136*cr1))
	b = uint8(clamp(y1 + 1.772*cb1))

	return r, g, b
}

func clamp(v float64) float64 {
	return math.Max(0, math.Min(255, v))
}
