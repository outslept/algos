package gausselimination

import "math"

func gaussElimination(a [][]float64, b []float64) []float64 {
	n := len(a)
	for i := 0; i < n; i++ {
		maxEl := math.Abs(a[i][i])
		maxRow := i
		for k := i + 1; k < n; k++ {
			if math.Abs(a[k][i]) > maxEl {
				maxEl = math.Abs(a[k][i])
				maxRow = k
			}
		}

		a[i], a[maxRow] = a[maxRow], a[i]
		b[i], b[maxRow] = b[maxRow], b[i]

		for k := i + 1; k < n; k++ {
			c := -a[k][i] / a[i][i]
			for j := i; j < n; j++ {
				if i == j {
					a[k][j] = 0
				} else {
					a[k][j] += c * a[i][j]
				}
			}
			b[k] += c * b[i]
		}
	}

	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		x[i] = b[i]
		for j := i + 1; j < n; j++ {
			x[i] -= a[i][j] * x[j]
		}
		x[i] /= a[i][i]
	}
	return x
}
