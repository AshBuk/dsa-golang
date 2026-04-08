// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// LeetCode 27: Remove Element
// https://leetcode.com/problems/remove-element/
//
// Given an integer array nums and an integer val, remove all occurrences of val
// in nums in-place. The order of the elements may be changed.
// Then return the number of elements in nums which are not equal to val.
//
// Example 1: Input: nums = [3,2,2,3], val = 3 Output: 2, nums = [2,2,_,_]
// Example 2: Input: nums = [0,1,2,2,3,0,4,2], val = 2 Output: 5, nums = [0,1,4,0,3,_,_,_]
//
// Constraints:
//   - 0 <= nums.length <= 100
//   - 0 <= nums[i] <= 50
//   - 0 <= val <= 100
//
// Time: O(n) - single pass with two pointers
// Space: O(1) - in-place

package main

import "fmt"

// 27: remove all occurrences of val in-place
func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)) // 5
}

/*
Slow/Fast Pointers pattern:
  - slow marks the write position, fast scans ahead
  - fast always moves forward; slow only advances on a valid write
  - everything before slow is the "clean" result

Remove val=2: nums = [0, 1, 2, 2, 3, 0, 4, 2]

  fast=0: 0!=2 → write nums[0]=0, slow++
          S
          F
        [ 0 | 1 | 2 | 2 | 3 | 0 | 4 | 2 ]

  fast=1: 1!=2 → write nums[1]=1, slow++
              S
              F
        [ 0 | 1 | 2 | 2 | 3 | 0 | 4 | 2 ]

  fast=2: 2==2 → skip
              S
                  F
        [ 0 | 1 | 2 | 2 | 3 | 0 | 4 | 2 ]

  fast=3: 2==2 → skip
              S
                      F
        [ 0 | 1 | 2 | 2 | 3 | 0 | 4 | 2 ]

  fast=4: 3!=2 → write nums[2]=3, slow++
                  S
                          F
        [ 0 | 1 | 3 | 2 | 3 | 0 | 4 | 2 ]

  fast=5: 0!=2 → write nums[3]=0, slow++
                      S
                              F
        [ 0 | 1 | 3 | 0 | 3 | 0 | 4 | 2 ]

  fast=6: 4!=2 → write nums[4]=4, slow++
                          S
                                  F
        [ 0 | 1 | 3 | 0 | 4 | 0 | 4 | 2 ]

  Return slow=5 → first 5 elements are the result

Key: overwrite (nums[slow] = nums[fast]) — tail values don't matter
*/
