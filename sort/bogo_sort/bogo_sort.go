package bogosort

import (
	"math/rand"
	"time"
)


/*
Bogosort, also known as permutation sort, stupid sort, slowsort, shotgun sort or monkey sort,
is a highly inefficient sorting algorithm based on the generate and test paradigm.

The algorithm successively generates permutations of its input until it finds one that is sorted.

See [Wikipedia](https://en.wikipedia.org/wiki/Bogosort)

Visual representation of the process:

    (3)---(1)---(4)---(2)
     |     |     |     |
     v     v     v     v
    (2)---(4)---(1)---(3)
     |     |     |     |
     v     v     v     v
    (1)---(2)---(3)---(4)

    1. Check if the array is sorted
    2. If not, randomly shuffle the entire array
    3. Repeat until sorted
*/


func bogoSort(arr []int) []int {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Keep shuffling until the array is sorted
	for !isSorted(arr) {
		shuffle(arr)
	}
	return arr
}

// isSorted checks if the array is in ascending order
func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i - 1] {
			return false
		}
	}
	return true
}

// shuffle randomly reorders elements in the array
func shuffle(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		// Generate a random index j such that 0 ≤ j ≤ i
		j := rand.Intn(i + 1)
		// Swap elements at i and j
		arr[i], arr[j] = arr[j], arr[i]
	}
}

/*
Time Complexity:
- Best Case: O(n), where n is the number of elements (when the array is already sorted)
- Worst Case: Unbounded (may never terminate)
- Average Case: O(n * n!), as there are n! permutations of n elements

Space Complexity: O(1) as it sorts in-place

Note: Bogosort is not used for practical sorting. It's primarily of interest
in theoretical computer science and as a humorous example of an inefficient algorithm.
In the worst case, it may never terminate if the correct permutation is never randomly selected.
*/
