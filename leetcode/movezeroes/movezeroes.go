// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// LeetCode 283: Move Zeroes
// https://leetcode.com/problems/move-zeroes/
//
// Given an integer array nums, move all 0's to the end of it
// while maintaining the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.
//
// Example 1: Input: nums = [0,1,0,3,12] Output: [1,3,12,0,0]
// Example 2: Input: nums = [0] Output: [0]
//
// Constraints:
//   - 1 <= nums.length <= 10^4
//   - -2^31 <= nums[i] <= 2^31 - 1
//
// Follow up: minimize the total number of operations done.
//
// Time: O(n) - single pass with two pointers
// Space: O(1) - in-place

package main

import "fmt"

// 283: move all zeroes to end, keeping order
func moveZeroes(nums []int) []int {
	slow, fast := 0, 0
	for slow < len(nums) && fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
	return nums
}

func main() {
	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12})) // [1 3 12 0 0]
}

/*
Slow/Fast Pointers pattern:
  - slow marks the write position, fast scans ahead
  - fast always moves forward; slow only advances on a valid write
  - everything before slow is the "clean" result

Move zeroes: nums = [0, 1, 0, 3, 12]

  fast=0: 0==0 → skip
          S
          F
        [ 0 | 1 | 0 | 3 | 12 ]

  fast=1: 1!=0 → swap(0,1), slow++
          S
              F
        [ 1 | 0 | 0 | 3 | 12 ]

  fast=2: 0==0 → skip
              S
                  F
        [ 1 | 0 | 0 | 3 | 12 ]

  fast=3: 3!=0 → swap(0,3), slow++
              S
                      F
        [ 1 | 3 | 0 | 0 | 12 ]

  fast=4: 12!=0 → swap(0,12), slow++
                  S
                          F
        [ 1 | 3 | 12 | 0 | 0 ]

  Return [1, 3, 12, 0, 0]

Key: swap (nums[slow], nums[fast]) — zeroes must end up at the end
*/
