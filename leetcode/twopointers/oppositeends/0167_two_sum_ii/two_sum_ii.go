// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// LeetCode 167: Two Sum II - Input Array Is Sorted
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
//
// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order,
// find two numbers such that they add up to a specific target number.
//
// Return the indices of the two numbers (1-indexed) as an integer array answer of size 2.
//
// Time: O(n) - each pointer moves at most n times total
// Space: O(1) - only two pointers used

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 26)) // [3 4]
}

/*
Two Pointers on sorted array: nums = [2, 7, 11, 15], target = 26

  Step 1:  sum = 2+15 = 17 < 26 → move left++
           L             R
         [ 2 | 7 | 11 | 15 ]

  Step 2:  sum = 7+15 = 22 < 26 → move left++
               L         R
         [ 2 | 7 | 11 | 15 ]

  Step 3:  sum = 11+15 = 26 == target → return [3, 4]
                   L     R
         [ 2 | 7 | 11 | 15 ]

Why it works:
  - Array is sorted, so left++ increases the sum, right-- decreases it
  - Each step eliminates one position → guaranteed O(n) convergence
  - No extra memory needed → O(1) space (vs O(n) for hash map approach)
*/
