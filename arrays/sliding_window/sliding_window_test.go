package slidingwindow

import "testing"

func TestMaxSumSlidingWindow(t *testing.T) {
    tests := []struct {
        name     string
        nums     []int
        k        int
        expected int
    }{
        {
            name:     "basic case",
            nums:     []int{1, 4, 2, 10, 2, 3, 1, 0, 20},
            k:        4,
            expected: 24,
        },
        {
            name:     "k equals array length",
            nums:     []int{1, 2, 3},
            k:        3,
            expected: 6,
        },
        {
            name:     "k greater than array length",
            nums:     []int{1, 2},
            k:        3,
            expected: 0,
        },
    }

    for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
					result := MaxSumSlidingWindow(tt.nums, tt.k)
					if result != tt.expected {
							t.Errorf("MaxSumSlidingWindow() = %v, want %v", result, tt.expected)
					}
			})
	}
}

func BenchmarkMaxSumSlidingWindow(b *testing.B) {
	nums := []int{1, 4, 2, 10, 2, 3, 1, 0, 20}
	k := 4
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
			MaxSumSlidingWindow(nums, k)
	}
}
