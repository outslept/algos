package rgb_to_lab

import "math"


func RGBToLAB(r, g, b uint8) (l, a, b2 float64) {
    x, y, z := RGBToXYZ(r, g, b)
    return XYZToLAB(x, y, z)
}

func XYZToLAB(x, y, z float64) (l, a, b float64) {
    x /= 0.95047
    y /= 1.00000
    z /= 1.08883

    x = labF(x)
    y = labF(y)
    z = labF(z)

    l = 116*y - 16
    a = 500 * (x - y)
    b = 200 * (y - z)

    return l, a, b
}

func labF(t float64) float64 {
	if t > 0.008856 {
			return math.Pow(t, 1.0/3.0)
	}
	return 7.787*t + 16.0/116.0
}
