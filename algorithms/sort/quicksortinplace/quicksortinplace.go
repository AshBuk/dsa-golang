// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// QuickSort In-Place Algorithm (Lomuto partition scheme)
// Time: O(n log n) average, O(n²) worst case
// Space: O(log n) - recursion stack only
// Note: modifies original array

package main

import "fmt"

func QuickSortInPlace(arr []int) {
	if len(arr) < 2 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, lo, hi int) {
	if lo < hi {
		pivotIdx := partition(arr, lo, hi)
		quickSort(arr, lo, pivotIdx-1)
		quickSort(arr, pivotIdx+1, hi)
	}
}

// Lomuto partition scheme
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
	arr := []int{4, 6, 2, 8, 3, 10, 1, 9, 5, 7}
	QuickSortInPlace(arr)
	fmt.Println(arr)
}

/*
How in-place quick sort works (Lomuto partition):

Example: [4, 6, 2, 8, 3, 10, 1, 9, 5, 7]

PARTITION (pivot = 7, last element):
    [4, 6, 2, 8, 3, 10, 1, 9, 5, 7]
                                 ↑ pivot = 7

    Scan & swap elements ≤ 7 to the left (storeIdx starts at -1):
    Note: storeIdx++ happens BEFORE swap
    i=0: 4≤7 → storeIdx(++)=0, swap arr[0]↔arr[0]  → [4, 6, 2, 8, 3, 10, 1, 9, 5, 7]
    i=1: 6≤7 → storeIdx=1, swap arr[1]↔arr[1]  → [4, 6, 2, 8, 3, 10, 1, 9, 5, 7]
    i=2: 2≤7 → storeIdx=2, swap arr[2]↔arr[2]  → [4, 6, 2, 8, 3, 10, 1, 9, 5, 7]
    i=3: 8>7 → skip, storeIdx stays at 2  ← storeIdx starts lagging behind i
    i=4: 3≤7 → storeIdx=3, swap arr[3]↔arr[4]  → [4, 6, 2, 3, 8, 10, 1, 9, 5, 7]
    i=5: 10>7 → skip
    i=6: 1≤7 → storeIdx=4, swap arr[4]↔arr[6]  → [4, 6, 2, 3, 1, 10, 8, 9, 5, 7]
    i=7: 9>7 → skip
    i=8: 5≤7 → storeIdx=5, swap arr[5]↔arr[8]  → [4, 6, 2, 3, 1, 5, 8, 9, 10, 7]

    Place pivot: swap arr[6]↔arr[9]
    Result: [4, 6, 2, 3, 1, 5, 7, 9, 10, 8]
            └─────≤7────────┘  ↑  └──>7───┘
                          pivotIdx=6

    Recursion: quickSort(0,5) and quickSort(7,9)

LEFT [4, 6, 2, 3, 1, 5] pivot = 5:
    After partition: [4, 2, 3, 1, 5, 6]
                     └───≤5────┘  ↑  >5
                              pivotIdx=4

    LEFT [4, 2, 3, 1] pivot = 1:
        After partition: [1, 2, 3, 4]
                          ↑  └>1───┘
                      pivotIdx=0

    Already sorted → done

RIGHT [9, 10, 8] pivot = 8:
    After partition: [8, 10, 9]
                      ↑  └>8─┘
                  pivotIdx=7

    [10, 9] pivot = 9:
        After partition: [9, 10]
                          ↑  >9
                      pivotIdx=8

RESULT: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] ✓

Key points:
- Pivot selection: LAST element (Lomuto scheme)
- In-place swaps: elements ≤ pivot move left of partition boundary
*/
