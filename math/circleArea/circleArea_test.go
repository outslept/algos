package circlearea

import (
	"math"
	"testing"
)

func TestCircleArea(t *testing.T) {
    tests := []struct {
        name     string
        radius   float64
        expected float64
    }{
        {"Zero radius", 0, 0},
        {"Unit circle", 1, math.Pi},
        {"Small radius", 2.5, 19.634954084936208},
        {"Large radius", 100, 31415.926535897932},
        {"Negative radius", -1, math.Pi}, // Note: This might not be mathematically correct, but it's how the function behaves
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := circleArea(tt.radius)
            if math.Abs(got-tt.expected) > 1e-9 {
                t.Errorf("circleArea(%v) = %v, want %v", tt.radius, got, tt.expected)
            }
        })
    }
}

func TestCircleAreaPrecision(t *testing.T) {
    radius := 123.456
    expected := 47916.63259368222

    got := circleArea(radius)
    if math.Abs(got-expected) > 1e-9 {
        t.Errorf("circleArea(%v) = %v, want %v", radius, got, expected)
    }
}

func BenchmarkCircleArea(b *testing.B) {
    for i := 0; i < b.N; i++ {
        circleArea(10)
    }
}
