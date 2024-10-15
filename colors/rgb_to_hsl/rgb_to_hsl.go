packge rgb_to_hsl

import "math"

func RGBToHSL(r, g, b uint 8) (h, s, l uint8) {
	r1 := float64(r) / 255
	g1 := float64(g) / 255
	b1 := float64(b) / 255

	max := math.Max(r1, math.Max(g1, b1))
	min := math.Min(r1, math.Min(g1, b1))

	l = (max + min) / 2

	if max == min {
			h = 0
			s = 0
	} else {
			d := max - min
			if l > 0.5 {
					s = d / (2 - max - min)
			} else {
					s = d / (max + min)
			}

			switch max {
			case r1:
					h = (g1 - b1) / d
					if g1 < b1 {
							h += 6
					}
			case g1:
					h = (b1 - r1) / d + 2
			case b1:
					h = (r1 - g1) / d + 4
			}
			h /= 6
	}

	return h * 360, s * 100, l * 100
}
