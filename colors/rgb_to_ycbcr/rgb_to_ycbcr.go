package rgb_to_ycbcr

func RGBToYCbCr(r, g, b uint8) (y, cb, cr uint8) {
	r1 := float64(r)
	g1 := float64(g)
	b1 := float64(b)

	y = uint8(0.299*r1 + 0.587*g1 + 0.114*b1)
	cb = uint8(128 - 0.168736*r1 - 0.331264*g1 + 0.5*b1)
	cr = uint8(128 + 0.5*r1 - 0.418688*g1 - 0.081312*b1)

	return y, cb, cr
}
