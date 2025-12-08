// QuickSelect Algorithm (In-Place)
// Time: O(n) average, O(n²) worst case
// Space: O(log n) - recursion stack only
// Purpose: finds k-th smallest element WITHOUT full sorting

package main

import (
	"fmt"
)

func quickSelect(arr []int, k int) int {
	lo := 0
	hi := len(arr) - 1

	for lo <= hi {
		// Lomuto partition
		pivot := arr[hi]
		storeIdx := lo - 1
		for i := lo; i < hi; i++ {
			if arr[i] <= pivot {
				storeIdx++
				arr[storeIdx], arr[i] = arr[i], arr[storeIdx]
			}
		}
		arr[storeIdx+1], arr[hi] = arr[hi], arr[storeIdx+1]
		pivotIdx := storeIdx + 1

		// Found k-th element
		if k-1 == pivotIdx {
			return arr[pivotIdx]
		}
		// Narrow search range
		if k-1 < pivotIdx {
			hi = pivotIdx - 1
		} else {
			lo = pivotIdx + 1
		}
	}
	return arr[lo]
}

func main() {
	arr := []int{4, 6, 2, 8, 3, 7, 1, 9, 5, 10}
	k := 4
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("%d-th smallest element: %d\n", k, quickSelect(arr, k))
}

/*
How QuickSelect (In-Place) works:
Example: [4, 6, 2, 8, 3, 7, 1, 9, 5, 10], k = 4 (find 4th smallest)

Iteration 1: lo=0, hi=9
    pivot = arr[9] = 10
    After partition: [4, 6, 2, 8, 3, 7, 1, 9, 5, 10]
                                            ↑ pivotIdx=9
    k-1=3, pivotIdx=9 → 3 < 9
    → Update hi = 8

Iteration 2: lo=0, hi=8
    pivot = arr[8] = 5
    After partition: [4, 2, 3, 1, 5, 7, 6, 9, 8]
                         ↑ pivotIdx=4
    k-1=3, pivotIdx=4 → 3 < 4
    → Update hi = 3

Iteration 3: lo=0, hi=3
    pivot = arr[3] = 1
    After partition: [1, 2, 3, 4]
                      ↑ pivotIdx=0
    k-1=3, pivotIdx=0 → 3 > 0
    → Update lo = 1

Iteration 4: lo=1, hi=3
    pivot = arr[3] = 4
    After partition: [1, 2, 3, 4]
                         ↑ pivotIdx=3
    k-1=3, pivotIdx=3 → 3 == 3
    → return arr[3] = 4 ✓

Algorithm (Lomuto partition):
    - pivot = arr[hi]
    - Partition: all ≤pivot go left
    - Compare k-1 with pivotIdx
    - Narrow search to [lo, pivotIdx-1] or [pivotIdx+1, hi]

vs QuickSort:
┌──────────────┬────────────────────────────┐
│ QuickSort    │ QuickSelect                │
├──────────────┼────────────────────────────┤
│ Both sides   │ One side only              │
│ O(n log n)   │ O(n) average               │
│ Sorts all    │ Finds k-th element         │
└──────────────┴────────────────────────────┘
*/
