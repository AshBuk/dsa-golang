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
// Time: O(n) - single pass with read/write pointers
// Space: O(1) - in-place

package main

import "fmt"

// 27: remove all occurrences of val in-place
func removeElement(nums []int, val int) int {
	write := 0
	for read := 0; read < len(nums); read++ {
		if nums[read] != val {
			nums[write] = nums[read]
			write++
		}
	}
	return write
}

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)) // 5
}

/*
Read/Write Pointers pattern:
  - write marks the next position to fill, read scans ahead
  - read always moves forward; write only advances on a valid write
  - everything before write is the "clean" result

Remove val=2: nums = [0, 1, 2, 2, 3, 0, 4, 2]

  read=0: 0!=2 → write nums[0]=0, write++
          W
          R
        [ 0 | 1 | 2 | 2 | 3 | 0 | 4 | 2 ]

  read=1: 1!=2 → write nums[1]=1, write++
              W
              R
        [ 0 | 1 | 2 | 2 | 3 | 0 | 4 | 2 ]

  read=2: 2==2 → skip
              W
                  R
        [ 0 | 1 | 2 | 2 | 3 | 0 | 4 | 2 ]

  read=3: 2==2 → skip
              W
                      R
        [ 0 | 1 | 2 | 2 | 3 | 0 | 4 | 2 ]

  read=4: 3!=2 → write nums[2]=3, write++
                  W
                          R
        [ 0 | 1 | 3 | 2 | 3 | 0 | 4 | 2 ]

  read=5: 0!=2 → write nums[3]=0, write++
                      W
                              R
        [ 0 | 1 | 3 | 0 | 3 | 0 | 4 | 2 ]

  read=6: 4!=2 → write nums[4]=4, write++
                          W
                                  R
        [ 0 | 1 | 3 | 0 | 4 | 0 | 4 | 2 ]

  Return write=5 → first 5 elements are the result

Key: overwrite (nums[write] = nums[read]) — tail values don't matter
*/
