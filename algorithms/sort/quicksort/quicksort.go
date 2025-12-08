// QuickSort Algorithm
// Time: O(n log n) average, O(n²) worst case
// Space: O(n) - creates new slices for left/right partitions
// Note: functional approach, does not modify original array

package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left := []int{}
	right := []int{}
	pivot := arr[0]

	for _, v := range arr[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
	fmt.Println(quickSort([]int{4, 6, 2, 8, 3, 7, 1, 9, 5, 10}))
}

/*
How quick sort works:

Example: [4, 6, 2, 8, 3, 7, 1, 9, 5, 10]

PARTITION PHASE (divide around pivot):
    [4, 6, 2, 8, 3, 7, 1, 9, 5, 10]
     ↑ pivot = 4
    left = [2,3,1]         (values <= 4)
    right = [6,8,7,9,5,10] (values > 4)

    Result: quickSort([2,3,1]) + [4] + quickSort([6,8,7,9,5,10])

RECURSION on left side:
    [2, 3, 1]
     ↑ pivot = 2
    left = [1]    (values <= 2)
    right = [3]   (values > 2)

    Result: quickSort([1]) + [2] + quickSort([3])
            → [1] + [2] + [3] → [1,2,3] ✓

RECURSION on right side:
    [6, 8, 7, 9, 5, 10]
     ↑ pivot = 6
    left = [5]            (values <= 6)
    right = [8,7,9,10]    (values > 6)

    Result: quickSort([5]) + [6] + quickSort([8,7,9,10])

    [8, 7, 9, 10]
     ↑ pivot = 8
    left = [7]     (values <= 8)
    right = [9,10] (values > 8)

    Result: [7] + [8] + quickSort([9,10])

    [9, 10]
     ↑ pivot = 9
    left = []      (values <= 9)
    right = [10]   (values > 9)

    Result: [] + [9] + [10] → [9,10] ✓

CONQUER PHASE (assembling sorted parts):
    [1,2,3] + [4] + [5,6,7,8,9,10] → [1,2,3,4,5,6,7,8,9,10] ✓

Key points:
- Pivot selection: first element (WARNING: causes O(n²) on sorted/reversed arrays)
- Elements are partitioned around pivot (<=pivot go left, >pivot go right)
- Sorting happens during partitioning and recursive assembly
- Each recursion places pivot in final sorted position
*/
