package karatsuba

import (
	"math"
	"strconv"
)

func karatsuba(x, y int64) int64 {
	if x < 10 || y < 10 {
		return x * y
	}

	n := max(len(strconv.FormatInt(x, 10)), len(strconv.FormatInt(y, 10)))
	m := (n + 1) / 2

	power := int64(math.Pow10(m))
	high1, low1 := x/power, x%power
	high2, low2 := y/power, y%power

	z0 := karatsuba(low1, low2)
	z1 := karatsuba((low1 + high1), (low2 + high2))
	z2 := karatsuba(high1, high2)

	return z2*power*power + (z1-z2-z0)*power + z0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
