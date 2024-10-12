package lcm

import (
	"testing"
)

func TestLCM(t *testing.T) {
        tests := []struct {
                name     string
                a        int
                b        int
                expected int
        }{
                {"LCM of 4 and 6", 4, 6, 12},
                {"LCM of 21 and 6", 21, 6, 42},
                {"LCM of 48 and 18", 48, 18, 144},
                {"LCM of prime numbers", 17, 19, 323},
                {"LCM with common factor", 24, 36, 72},
                {"LCM with one number as factor of other", 5, 15, 15},
                {"LCM of same numbers", 7, 7, 7},
                {"LCM with 1", 1, 5, 5},
                {"LCM with 0", 0, 5, 0},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        result := lcm(tt.a, tt.b)
                        if result != tt.expected {
                                t.Errorf("lcm(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
                        }
                })
        }
}

func TestGCD(t *testing.T) {
        tests := []struct {
                name     string
                a        int
                b        int
                expected int
        }{
                {"GCD of 48 and 18", 48, 18, 6},
                {"GCD of 54 and 24", 54, 24, 6},
                {"GCD of prime numbers", 17, 19, 1},
                {"GCD with common factor", 24, 36, 12},
                {"GCD of same numbers", 7, 7, 7},
                {"GCD with 1", 1, 5, 1},
                {"GCD with 0", 0, 5, 5},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        result := gcd(tt.a, tt.b)
                        if result != tt.expected {
                                t.Errorf("gcd(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
                        }
                })
        }
}

func TestLCMCommutative(t *testing.T) {
        a, b := 24, 36
        result1 := lcm(a, b)
        result2 := lcm(b, a)
        if result1 != result2 {
                t.Errorf("lcm(%d, %d) = %d, but lcm(%d, %d) = %d", a, b, result1, b, a, result2)
        }
}

func TestGCDCommutative(t *testing.T) {
        a, b := 48, 18
        result1 := gcd(a, b)
        result2 := gcd(b, a)
        if result1 != result2 {
                t.Errorf("gcd(%d, %d) = %d, but gcd(%d, %d) = %d", a, b, result1, b, a, result2)
        }
}

func BenchmarkLCM(b *testing.B) {
        for i := 0; i < b.N; i++ {
                lcm(48, 18)
        }
}

func BenchmarkGCD(b *testing.B) {
        for i := 0; i < b.N; i++ {
                gcd(48, 18)
        }
}
