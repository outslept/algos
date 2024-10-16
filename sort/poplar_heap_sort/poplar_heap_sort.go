package poplarheapsort

/*
Poplar Heap Sort is a variation of the Heap Sort algorithm that uses a binary heap data structure.
It first builds a max heap from the input array and then repeatedly extracts the maximum element
from the heap and places it at the end of the sorted portion of the array.

See [Wikipedia](https://en.wikipedia.org/wiki/Heapsort) for general Heap Sort information.

Visual representation of the process:

    Initial array:     [4, 10, 3, 5, 1]

    Build max heap:    [10, 5, 3, 4, 1]
                         /\
                        /  \
                       /    \
                      /      \
                     5        3
                    / \      /
                   4   1

    Extract max:      [1, 5, 3, 4] | [10]
                       /\
                      /  \
                     /    \
                    5      3
                   / \
                  4   1

    Final sorted:     [1, 3, 4, 5, 10]
*/

func poplarHeapSort(arr []int) []int {
        n := len(arr)
        // Build a max heap
        buildPoplarHeap(arr)

        // One by one extract elements from the heap
        for i := n - 1; i > 0; i-- {
                // Move current root to end
                arr[0], arr[i] = arr[i], arr[0]
                // Call siftDown on the reduced heap
                siftDown(arr, 0, i)
        }
        return arr
}

// buildPoplarHeap builds a max heap from an unsorted array
func buildPoplarHeap(arr []int) {
        // Start from the last non-leaf node and heapify all nodes
        for i := len(arr)/2 - 1; i >= 0; i-- {
                siftDown(arr, i, len(arr))
        }
}

// siftDown maintains the max heap property
func siftDown(arr []int, root, n int) {
        for {
                // Find the left child of the current root
                child := 2*root + 1
                // If the child is out of bounds, we're done
                if child >= n {
                        break
                }
                // If there's a right child and it's larger, use it instead
                if child+1 < n && arr[child] < arr[child+1] {
                        child++
                }
                // If the root is already larger than the largest child, we're done
                if arr[root] >= arr[child] {
                        break
                }
                // Otherwise, swap the root with the largest child
                arr[root], arr[child] = arr[child], arr[root]
                // Continue sifting down from the child's position
                root = child
        }
}

/*
Time Complexity:
- Best Case: O(n log n)
- Average Case: O(n log n)
- Worst Case: O(n log n)

Space Complexity: O(1) (in-place algorithm)
*/
