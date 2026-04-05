// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// Merge Sort Algorithm
// Time: O(n log n) - all cases (worst, average, best)
// Space: O(n) - creates temporary arrays during merging

package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := []int{}
	var i, j int

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func main() {
	fmt.Println(mergeSort([]int{1, 2, 4, 3, 6, 5, 9, 8, 7}))
}

/*
Call Stack:

mergeSort([1,2,4,3,6,5,9,8,7])
├─ mergeSort([1,2,4,3])
│  ├─ mergeSort([1,2]) -> merge([1],[2]) -> [1,2]
│  ├─ mergeSort([4,3]) -> merge([4],[3]) -> [3,4]
│  └─ merge([1,2],[3,4]) -> [1,2,3,4]
└─ mergeSort([6,5,9,8,7])
   ├─ mergeSort([6,5]) -> merge([6],[5]) -> [5,6]
   ├─ mergeSort([9,8,7])
   │  ├─ mergeSort([9]) -> [9]
   │  └─ mergeSort([8,7]) -> merge([8],[7]) -> [7,8]
   │  └─ merge([9],[7,8]) -> [7,8,9]
   └─ merge([5,6],[7,8,9]) -> [5,6,7,8,9]
└─ merge([1,2,3,4],[5,6,7,8,9]) -> final [1..9]

How merge sort works:

Example: [1, 2, 4, 3, 6, 5, 9, 8, 7]

DIVIDE PHASE (going down recursion):  // mid := len(arr) / 2
        [1, 2, 4, 3, 6, 5, 9, 8, 7]
              /                \
      [1, 2, 4, 3]          [6, 5, 9, 8, 7]
         /      \              /          \
     [1, 2]   [4, 3]      [6, 5]      [9, 8, 7]
      / \      / \         / \          /     \
    [1] [2]  [4] [3]     [6] [5]     [9]    [8, 7]
                                              / \
                                            [8] [7]
    ↑ BASE CASE - 9 slices of size 1  // if len(arr) < 2

CONQUER PHASE (going up, merge sorts):
    Level 1: merge pairs
        [1] + [2] → [1, 2]
        [4] + [3] → [3, 4]
        [6] + [5] → [5, 6]
        [8] + [7] → [7, 8]

    Level 2: merge groups
        [1, 2] + [3, 4] → [1, 2, 3, 4]
        [5, 6] + [9] → [5, 6, 9]
        [9] + [7, 8] → [7, 8, 9]

    Level 3: merge larger groups
        [1, 2, 3, 4] + [5, 6, 9] → [1, 2, 3, 4, 5, 6, 9]

    Level 4: final merge
        [1, 2, 3, 4] + [5, 6, 7, 8, 9] → [1, 2, 3, 4, 5, 6, 7, 8, 9] ✓

Key: mergeSort() decomposes, merge() sorts while combining
Time: O(n log n), Space: O(n)
*/
