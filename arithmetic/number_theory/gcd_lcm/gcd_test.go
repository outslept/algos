package gcd

import (
        "testing"
				"fmt"
)

func TestGCD(t *testing.T) {
        tests := []struct {
                a, b     int
                expected int
        }{
                {48, 18, 6},
                {18, 48, 6},
                {0, 0, 0},
                {0, 5, 5},
                {5, 0, 5},
                {15, 5, 5},
                {5, 15, 5},
                {17, 13, 1},
        }

        for _, test := range tests {
                t.Run(fmt.Sprintf("gcd(%d, %d)", test.a, test.b), func(t *testing.T) {
                        result := gcd(test.a, test.b)
                        if result != test.expected {
                                t.Errorf("Expected gcd(%d, %d) = %d, but got %d", test.a, test.b, test.expected, result)
                        }
                })
        }
}

func BenchmarkGCD(b *testing.B) {
        for i := 0; i < b.N; i++ {
                gcd(123456, 7890)
        }
}
