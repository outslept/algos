package colorconv

import "math"

func HSVToRGB(h, s, v float64) (r, g, b uint8) {
    h /= 60
    s /= 100
    v /= 100

    hi := math.Floor(h)
    f := h - hi
    p := v * (1 - s)
    q := v * (1 - s*f)
    t := v * (1 - s*(1-f))

    var r1, g1, b1 float64

    switch int(hi) {
    case 0, 6:
        r1, g1, b1 = v, t, p
    case 1:
        r1, g1, b1 = q, v, p
    case 2:
        r1, g1, b1 = p, v, t
    case 3:
        r1, g1, b1 = p, q, v
    case 4:
        r1, g1, b1 = t, p, v
    case 5:
        r1, g1, b1 = v, p, q
    }

    return uint8(r1 * 255), uint8(g1 * 255), uint8(b1 * 255)
}
