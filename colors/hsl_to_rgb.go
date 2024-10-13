// hsl_to_rgb.go
// description: Implementation of HSL to RGB color conversion.
// details:
// This file contains a function to convert colors from the HSL (Hue, Saturation, Lightness)
// color model to to the RGB (Red, Green, Blue) color model.
// See https://en.wikipedia.org/wiki/HSL_and_HSV for more information on the HSL color model.

package hsl_to_rgb

import "math"

// HSLToRGB converts a color from HSL to RGB color space.
// It takes three uint8 values representing HSL (0-360, 0-100, 0-100) and returns three uint8 values for RGB (0-255).
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

// hueToRGB is a helper function for HSLToRGB.
// It converts the hue to RGB value.
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
