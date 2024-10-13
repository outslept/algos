package rgb_to_hwb

import "math"

func RGBToHWB(r, g, b uint8) (h, w, bl float64) {
    r1 := float64(r) / 255
    g1 := float64(g) / 255
    b1 := float64(b) / 255

    max := math.Max(r1, math.Max(g1, b1))
    min := math.Min(r1, math.Min(g1, b1))

    if max == min {
        h = 0
    } else if max == r1 {
        h = (g1 - b1) / (max - min)
    } else if max == g1 {
        h = 2 + (b1-r1)/(max-min)
    } else {
        h = 4 + (r1-g1)/(max-min)
    }

    h *= 60
    if h < 0 {
        h += 360
    }

    w = min
    bl = 1 - max

    return h, w * 100, bl * 100
}
