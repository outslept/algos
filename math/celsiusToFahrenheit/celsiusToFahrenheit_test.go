package celsiustofahrenheit

import (
	"math"
	"testing"
)

func TestCelsiusToFahrenheit(t *testing.T) {
        tests := []struct {
                name     string
                celsius  float64
                expected float64
        }{
                {"Freezing point of water", 0, 32},
                {"Boiling point of water", 100, 212},
                {"Body temperature", 37, 98.6},
                {"Absolute zero", -273.15, -459.67},
                {"Negative temperature", -40, -40},
                {"Positive temperature", 25, 77},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        got := celsiusToFahrenheit(tt.celsius)
                        if math.Abs(got-tt.expected) > 1e-6 {
                                t.Errorf("celsiusToFahrenheit(%v) = %v, want %v", tt.celsius, got, tt.expected)
                        }
                })
        }
}

func TestCelsiusToFahrenheitPrecision(t *testing.T) {
        celsius := 36.5
        expected := 97.7

        got := celsiusToFahrenheit(celsius)
        if math.Abs(got-expected) > 1e-6 {
                t.Errorf("celsiusToFahrenheit(%v) = %v, want %v", celsius, got, expected)
        }
}

func BenchmarkCelsiusToFahrenheit(b *testing.B) {
        for i := 0; i < b.N; i++ {
                celsiusToFahrenheit(20)
        }
}
