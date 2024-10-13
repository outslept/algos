// cmyk_to_rgb.go
// description: Implementation of CMYK to RGB color conversion.
// details:
// This file contains a function to convert colors from the CMYK (Cyan, Magenta, Yellow, Key/Black)
// color model to to the RGB (Red, Green, Blue) color model.
// See https://en.wikipedia.org/wiki/CMYK_color_model for more information on the CMYK color model.
package cmyk_to_rgb

// CMYKToRGB converts a color from CMYK to RGB color space.
// It takes four uint8 values representing CMYK (0-255) and returns three uint8 values for RGB (0-255).
func CMYKToRGB(c, m, y, k uint8) (r, g, b uint8) {
	r = uint8((1 - c) * (1 - k) * 255)
	g = uint8((1 - m) * (1 - k) * 255)
	b = uint8((1 - y) * (1 - k) * 255)

	return r, g, b
}
