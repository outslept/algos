package hsl_to_rgb

import "math"

func HSLToRGB(h, s, l uint8) (r, g, b uint8) {
	h /= 360
	s /= 100
	l /= 100

	var r1, g1, b1 float64

	if s == 0 {
		r1 = l
		g1 = l
		b1 = l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q

		r1 = hueToRGB(p, q, h+1/3)
		g1 = hueToRGB(p, q, h)
		b1 = hueToRGB(p, q, h-1/3)
	}

	return uint8(r1 * 255), uint8(g1 * 255), uint8(b1 * 255)
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1/6 {
		return p + (q-p)*6*t
	}
	if t < 1/2 {
		return q
	}
	if t < 2/3 {
		return p + (q-p)*(2/3-t)*6
	}
	return p
}
