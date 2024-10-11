package estimatePi

import (
	"math"
	"testing"
)

func TestEstimatePi(t *testing.T) {
        tests := []struct {
                name       string
                iterations int
                tolerance  float64
        }{
                {"Low iterations", 1000, 0.1},
                {"Medium iterations", 100000, 0.01},
                {"High iterations", 1000000, 0.001},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        result := estimatePi(tt.iterations)
                        if math.Abs(result-math.Pi) > tt.tolerance {
                                t.Errorf("estimatePi(%d) = %v, want %v ± %v", tt.iterations, result, math.Pi, tt.tolerance)
                        }
                })
        }
}

func TestEstimatePiConsistency(t *testing.T) {
        iterations := 1000000
        runs := 5
        results := make([]float64, runs)

        for i := 0; i < runs; i++ {
                results[i] = estimatePi(iterations)
        }

        for i := 1; i < runs; i++ {
                if math.Abs(results[i]-results[0]) > 0.01 {
                        t.Errorf("Inconsistent results: run 1 = %v, run %d = %v", results[0], i+1, results[i])
                }
        }
}

func BenchmarkEstimatePi(b *testing.B) {
        for i := 0; i < b.N; i++ {
                estimatePi(100000)
        }
}
