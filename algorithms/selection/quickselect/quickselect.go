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
		// Lomuto partition: split array into [<=pivot] | pivot | [>=pivot]
		pivot := arr[hi]
		storeIdx := lo - 1 // boundary: elements at indices <=storeIdx are <=pivot

		for i := lo; i < hi; i++ {
			if arr[i] <= pivot {
				// Found "good" element (<=pivot), expand left zone
				// storeIdx+1 holds a "frozen" bad element (>pivot) - push it right
				storeIdx++
				arr[storeIdx], arr[i] = arr[i], arr[storeIdx]
			}
		}
		// Place pivot right after all elements ≤ it
		pivotIdx := storeIdx + 1
		arr[pivotIdx], arr[hi] = arr[hi], arr[pivotIdx]

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
[4 6 2 8 3 7 1 9 5 10], k = 4 (find 4th smallest)

lo=0, hi=9, pivot=10
[4 6 2 8 3 7 1 9 5 | 10]
                      pivotIdx=9
k-1 < pivotIdx → hi=8
----------------------------

lo=0, hi=8, pivot=5
[4 2 3 1 | 5 | 7 6 9 8 10]
            pivotIdx=4
k-1 < pivotIdx → hi=3
----------------------------

lo=0, hi=3, pivot=1
[1 | 2 3 4 | 5 7 6 9 8 10]
  pivotIdx=0
k-1 > pivotIdx → lo=1
----------------------------

lo=1, hi=3, pivot=4
[1 2 3 | 4 | 5 7 6 9 8 10]
          pivotIdx=3
k-1 == pivotIdx → return 4
----------------------------

Invariant:
left ≤ pivot | pivot | right > pivot


vs QuickSort:
┌──────────────┬────────────────────────────┐
│ QuickSort    │ QuickSelect                │
├──────────────┼────────────────────────────┤
│ Both sides   │ One side only              │
│ O(n log n)   │ O(n) average               │
│ Sorts all    │ Finds k-th element         │
└──────────────┴────────────────────────────┘
*/
