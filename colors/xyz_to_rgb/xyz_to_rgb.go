package xyz_to_rgb

import "math"

func XYZToRGB(x, y, z float64) (r, g, b uint8) {
    r1 := 3.2404542*x - 1.5371385*y - 0.4985314*z
    g1 := -0.9692660*x + 1.8760108*y + 0.0415560*z
    b1 := 0.0556434*x - 0.2040259*y + 1.0572252*z

    r = uint8(clamp(gammaInverse(r1) * 255))
    g = uint8(clamp(gammaInverse(g1) * 255))
    b = uint8(clamp(gammaInverse(b1) * 255))

    return r, g, b
}

func gammaInverse(c float64) float64 {
    if c > 0.0031308 {
        return 1.055*math.Pow(c, 1/2.4) - 0.055
    }
    return 12.92 * c
}

func clamp(v float64) float64 {
    return math.Max(0, math.Min(255, v))
}
