package fibonacci

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
        tests := []struct {
                name     string
                n        int
                expected int
        }{
                {"0th Fibonacci", 0, 0},
                {"1st Fibonacci", 1, 1},
                {"2nd Fibonacci", 2, 1},
                {"3rd Fibonacci", 3, 2},
                {"5th Fibonacci", 5, 5},
                {"10th Fibonacci", 10, 55},
                {"20th Fibonacci", 20, 6765},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        result := fibonacci(tt.n)
                        if result != tt.expected {
                                t.Errorf("fibonacci(%d) = %d; want %d", tt.n, result, tt.expected)
                        }
                })
        }
}

func TestFibonacciNegative(t *testing.T) {
        result := fibonacci(-1)
        if result != -1 {
                t.Errorf("fibonacci(-1) = %d; want -1", result)
        }
}

func TestFibonacciLarge(t *testing.T) {
        n := 40
        expected := 102334155
        result := fibonacci(n)
        if result != expected {
                t.Errorf("fibonacci(%d) = %d; want %d", n, result, expected)
        }
}

func BenchmarkFibonacci(b *testing.B) {
        for i := 0; i < b.N; i++ {
                fibonacci(20)
        }
}

func BenchmarkFibonacciLarge(b *testing.B) {
        for i := 0; i < b.N; i++ {
                fibonacci(40)
        }
}
