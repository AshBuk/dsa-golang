// Bubble Sort Algorithm
// Time Complexity: O(n²) - worst and average case, O(n) - best case (already sorted)
// Space Complexity: O(1) - sorts in-place

package main

import "fmt"

func bubbleSort(arr []int) []int {
	hi := len(arr) - 1
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < hi; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
		hi--
	}
	return arr
}

func main() {
	fmt.Println(bubbleSort([]int{1, 3, 4, 2, 7, 5, 6, 8}))
}

/*
How bubble sort works:

Example: [1, 3, 4, 2, 7, 5, 6, 8]

PASS 1 (compare adjacent pairs, swap if left > right):
    [1, 3, 4, 2, 7, 5, 6, 8]
     1≤3 ✓  3≤4 ✓  4>2 ✗ swap  2≤7 ✓  7>5 ✗ swap  5≤6 ✓  6≤8 ✓
    [1, 3, 2, 4, 5, 7, 6, 8]  ← continue comparing pairs
     1≤3 ✓  3>2 ✗ swap  2≤4 ✓  4≤5 ✓  5≤7 ✓  7>6 ✗ swap  6≤8 ✓
    [1, 2, 3, 4, 5, 6, 7, 8]  ← largest element (8) "bubbled up" to end

PASS 2 (hi = 6, last element already sorted):
    [1, 2, 3, 4, 5, 6, 7, 8]
     1≤2 ✓  2≤3 ✓  3≤4 ✓  4≤5 ✓  5≤6 ✓  6≤7 ✓
    [1, 2, 3, 4, 5, 6, 7, 8]  ← no swaps, sorted = true, exit early

Key points:
- Compares adjacent elements and swaps if left > right
- After each pass, largest unsorted element "bubbles up" to its final position
- hi decreases each pass (optimization: skip already sorted end)
- sorted flag allows early exit when array is already sorted
- Best case O(n) when array is already sorted (one pass with no swaps)
*/
