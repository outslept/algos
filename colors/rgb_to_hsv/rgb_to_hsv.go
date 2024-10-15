package rgb_to_hsv

import "math"

func RGBToHSV(r, g, b uint8) (h, s, v float64) {
    r1 := float64(r) / 255
    g1 := float64(g) / 255
    b1 := float64(b) / 255

    max := math.Max(r1, math.Max(g1, b1))
    min := math.Min(r1, math.Min(g1, b1))
    delta := max - min

    v = max

    if max != 0 {
        s = delta / max
    } else {
        s = 0
        h = -1
        return
    }

    if r1 == max {
        h = (g1 - b1) / delta
    } else if g1 == max {
        h = 2 + (b1-r1)/delta
    } else {
        h = 4 + (r1-g1)/delta
    }

    h *= 60
    if h < 0 {
        h += 360
    }

    return h, s * 100, v * 100
}
