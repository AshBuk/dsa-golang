// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// LeetCode 349: Intersection of Two Arrays
// https://leetcode.com/problems/intersection-of-two-arrays/
//
// Given two integer arrays nums1 and nums2, return an array of their intersection.
// Each element in the result must be unique. The result can be in any order.
//
// Time: O(m+n) - one pass per array
// Space: O(m) - set stores nums1 elements
//
// LeetCode 350: Intersection of Two Arrays II
// https://leetcode.com/problems/intersection-of-two-arrays-ii/
//
// Given two integer arrays nums1 and nums2, return an array of their intersection.
// Each element in the result must appear as many times as it shows in both arrays.
//
// Time: O(m+n) - one pass per array
// Space: O(m) - frequency map stores nums1 counts

package main

import "fmt"

// 349: unique intersection using set
func intersection(nums1, nums2 []int) []int {
	set := make(map[int]bool)
	for _, num := range nums1 {
		set[num] = true
	}
	var result []int
	for _, num := range nums2 {
		if set[num] {
			result = append(result, num)
			delete(set, num) // avoid duplicates in result
		}
	}
	return result
}

// 350: intersection preserving duplicates using frequency map
func intersect(nums1, nums2 []int) []int {
	freq := make(map[int]int)
	for _, num := range nums1 {
		freq[num]++
	}
	var result []int
	for _, num := range nums2 {
		if freq[num] > 0 {
			result = append(result, num)
			freq[num]--
		}
	}
	return result
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))    // [2]
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8})) // [9 4]

	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))    // [2 2]
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8})) // [9 4]
}

/*
-#349 — Set approach: nums1 = [1,2,2,1], nums2 = [2,2]

  Build set from nums1:  {1, 2}
  Walk nums2:
    num=2, in set → add to result, delete from set → result=[2]
    num=2, not in set (deleted) → skip
  Return [2]

-#350 — Frequency map: nums1 = [1,2,2,1], nums2 = [2,2]

  Build freq from nums1:  {1:2, 2:2}
  Walk nums2:
    num=2, freq[2]=2 > 0 → add to result, freq[2]-- → result=[2], freq={1:2, 2:1}
    num=2, freq[2]=1 > 0 → add to result, freq[2]-- → result=[2,2], freq={1:2, 2:0}
  Return [2, 2]

Key difference:
  -#349: set (bool map) + delete → each value appears at most once
  -#350: frequency map (int map) + decrement → preserves duplicate count
*/
