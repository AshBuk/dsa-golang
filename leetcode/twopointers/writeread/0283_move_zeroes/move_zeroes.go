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
// Time: O(n) - single pass with read/write pointers
// Space: O(1) - in-place

package main

import "fmt"

// 283: move all zeroes to end, keeping order
func moveZeroes(nums []int) []int {
	write := 0
	for read := 0; read < len(nums); read++ {
		if nums[read] != 0 {
			nums[write], nums[read] = nums[read], nums[write]
			write++
		}
	}
	return nums
}

func main() {
	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12})) // [1 3 12 0 0]
}

/*
Read/Write Pointers pattern:
  - write marks the next position to fill, read scans ahead
  - read always moves forward; write only advances on a valid write
  - everything before write is the "clean" result

Move zeroes: nums = [0, 1, 0, 3, 12]

  read=0: 0==0 → skip
          W
          R
        [ 0 | 1 | 0 | 3 | 12 ]

  read=1: 1!=0 → swap(0,1), write++
          W
              R
        [ 1 | 0 | 0 | 3 | 12 ]

  read=2: 0==0 → skip
              W
                  R
        [ 1 | 0 | 0 | 3 | 12 ]

  read=3: 3!=0 → swap(0,3), write++
              W
                      R
        [ 1 | 3 | 0 | 0 | 12 ]

  read=4: 12!=0 → swap(0,12), write++
                  W
                          R
        [ 1 | 3 | 12 | 0 | 0 ]

  Return [1, 3, 12, 0, 0]

Key: swap (nums[write], nums[read]) — zeroes must end up at the end
*/
