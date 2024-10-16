package gnomesort

/*
Gnome Sort, also called Stupid Sort, is a sorting algorithm originally proposed
by Dr. Hamid Sarbazi-Azad in 2000. It is a simple sorting algorithm that is similar
to Insertion Sort, but moves elements to their proper position by a series of swaps,
similar to Bubble Sort.

The algorithm gets its name from the behavior of a garden gnome sorting his flower pots.

See [Wikipedia](https://en.wikipedia.org/wiki/Gnome_sort)

Visual representation of the process:

    (5)---(3)---(4)---(1)---(2)
     ^
     |
    (3)---(5)---(4)---(1)---(2)
          ^
          |
    (3)---(4)---(5)---(1)---(2)
               ^
               |
    (3)---(4)---(1)---(5)---(2)
          ^
          |
    (3)---(1)---(4)---(5)---(2)
     ^
     |
    (1)---(3)---(4)---(5)---(2)
          ^
          |
    ...

    1. Compare current element with previous
    2. If out of order, swap and move back
    3. If in order or at start, move forward
    4. Repeat until end of array is reached
*/

func gnomeSort(arr []int) []int {
    i := 1
    for i < len(arr) {
        if i == 0 || arr[i-1] <= arr[i] {
            // If at start or current element is in correct order,
            // move to next element
            i++
        } else {
            // If current element is smaller than previous,
            // swap them and move back one position
            arr[i], arr[i-1] = arr[i-1], arr[i]
            i--
        }
    }
    return arr
}

/*
Time Complexity:
- Worst Case: O(n^2), where n is the number of elements
- Best Case: O(n) when the array is already sorted
- Average Case: O(n^2)

Space Complexity: O(1) as it sorts in-place

Note: While simple to implement, Gnome Sort is not efficient for large datasets.
It's mainly used for educational purposes or when simplicity is more important
than efficiency. It performs better than Bubble Sort on arrays that are nearly sorted.
*/
