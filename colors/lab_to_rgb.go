package colorconv

import "math"

func LABToRGB(l, a, b float64) (r, g, b2 uint8) {
    x, y, z := LABToXYZ(l, a, b)
    return XYZToRGB(x, y, z)
}

func LABToXYZ(l, a, b float64) (x, y, z float64) {
    y = (l + 16) / 116
    x = a/500 + y
    z = y - b/200

    x = labInvF(x) * 0.95047
    y = labInvF(y) * 1.00000
    z = labInvF(z) * 1.08883

    return x, y, z
}

func labInvF(t float64) float64 {
    if t > 0.206893 {
        return math.Pow(t, 3)
    }
    return (t - 16.0/116.0) / 7.787
}
