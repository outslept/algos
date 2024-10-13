package rgb_to_lch

import "math"

func RGBToLCH(r, g, b uint8) (l, c, h float64) {
	lab_l, lab_a, lab_b := RGBToLAB(r, g, b)

	l = lab_l
	c = math.Sqrt(lab_a*lab_a + lab_b*lab_b)
	h = math.Atan2(lab_b, lab_a) * 180 / math.Pi

	if h < 0 {
			h += 360
	}

	return l, c, h
}
