package lch_to_rgb

import "math"

func LCHToRGB(l, c, h float64) (r, g, b uint8) {
    h = h * math.Pi / 180
    
    a := c * math.Cos(h)
    b2 := c * math.Sin(h)
    
    return LABToRGB(l, a, b2)
}
