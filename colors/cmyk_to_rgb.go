package cmyk_to_rgb

func CMYKToRGB(c, m, y, k uint8) (r, g, b uint8) {
	r = uint8((1 - c) * (1 - k) * 255)
	g = uint8((1 - m) * (1 - k) * 255)
	b = uint8((1 - y) * (1 - k) * 255)
	return r, g, b
}
