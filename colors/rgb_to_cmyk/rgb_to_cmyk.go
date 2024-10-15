package rgb_to_cmyk

func RGBToCMYK(r, g, b uint8) (c, m, y, k uint8) {
	if r == 0 && g == 0 && b == 0 {
		return 0, 0, 0, 1
	}

	r1 := float64(r) / 255
	g1 := float64(g) / 255
	b1 := float64(b) / 255

	k = 1 - max(r1, g1, b1)
	c = (1 - r1 - k) / (1 - k)
	m = (1 - g1 - k) / (1 - k)
	y = (1 - b1 - k) / (1 - k)

	return c, m, y, k
}

func max(a, b, c float64) float64 {
	if a > b && a > c {
		return a
	}
	if b > c {
		return b
	}
	return c
}
