// QuickSort In-Place Algorithm (Lomuto partition scheme)
// Time: O(n log n) average, O(n²) worst case
// Space: O(log n) - recursion stack only
// Note: modifies original array

package main

import "fmt"

func qSortInPlace(arr []int, lo, hi int) {
	if lo < hi {
		pivotIdx := partition(arr, lo, hi)
		qSortInPlace(arr, lo, pivotIdx-1) // sort left
		qSortInPlace(arr, pivotIdx+1, hi) // sort right
	}
}

// Lomuto partition scheme
// Places pivot element in correct position and returns its index
func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	storeIdx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			storeIdx++
			arr[storeIdx], arr[i] = arr[i], arr[storeIdx]
		}
	}
	arr[storeIdx+1], arr[hi] = arr[hi], arr[storeIdx+1]
	return storeIdx + 1
}

func main() {
	arr := []int{4, 6, 2, 8, 3, 7, 1, 9, 5, 10}
	lo := 0
	hi := len(arr) - 1
	qSortInPlace(arr, lo, hi)
	fmt.Println(arr)
}

/*
How in-place quick sort works (Lomuto partition):

Example: [4, 6, 2, 8, 3, 7, 1, 9, 5, 10]

PARTITION PHASE (in-place rearrangement):
    Initial: [4, 6, 2, 8, 3, 7, 1, 9, 5, 10]
                                         ↑ pivot = 10 (last element)

    After partition: [4, 6, 2, 8, 3, 7, 1, 9, 5, 10]
                     └──────elements ≤ 10──────┘ │
                     All elements fit, pivot stays at end
                     pivotIdx = 9

    Recursion: quickSort(arr, 0, 8) and quickSort(arr, 10, 9) - nothing

First recursion on [4, 6, 2, 8, 3, 7, 1, 9, 5]:
    pivot = 5 (last)
    After swaps: [4, 2, 3, 1, 5, 7, 6, 9, 8]
                 └──≤5─────┘  │  └──>5────┘
                 pivotIdx = 4

    Recursion: quickSort(arr, 0, 3) and quickSort(arr, 5, 8)

Continue left [4, 2, 3, 1]:
    pivot = 1
    After swaps: [1, 2, 3, 4]
                  │ └─>1───┘
                 pivotIdx = 0

Continue right [7, 6, 9, 8]:
    pivot = 8
    After swaps: [7, 6, 8, 9]
                 └≤8─┘  │  >8
                 pivotIdx = 2

... recursion continues until sorted

RESULT: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] ✓

Key points:
- Pivot selection: LAST element (Lomuto scheme)
- In-place swaps: elements ≤ pivot move left of partition boundary
*/
