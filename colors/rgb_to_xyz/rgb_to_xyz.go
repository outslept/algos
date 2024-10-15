package rgb_to_xyz

import "math"

func RGBToXYZ(r, g, b uint8) (x, y, z uint8) {
	r1 := gammaCorrect(float64(r) / 255)
	g1 := gammaCorrect(float64(g) / 255)
	b1 := gammaCorrect(float64(b) / 255)

	x = 0.4124564*r1 + 0.3575761*g1 + 0.1804375*b1
	y = 0.2126729*r1 + 0.7151522*g1 + 0.0721750*b1
	z = 0.0193339*r1 + 0.1191920*g1 + 0.9503041*b1

	return x, y, z
}

func gammaCorrect(c float64) float64 {
	if c > 0.04045 {
		return math.Pow((c + 0.055) / 1.055, 2.4)
	}
	return c / 12.92
}
