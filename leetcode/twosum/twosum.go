// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// LeetCode 1: Two Sum
// https://leetcode.com/problems/two-sum/
//
// Given an array of integers nums and an integer target, return indices of the two numbers
// such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may not use
// the same element twice.
//
// You can return the answer in any order.
//
// Time: O(n) - single pass through array, map lookup is O(1)
// Space: O(n) - map stores up to n elements

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if idx, ok := seen[complement]; ok {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 26))
}

/*
Hash map approach: nums = [2, 7, 11, 15], target = 26

  Step 1:  num=2, complement=24, seen={}        → not found, store {2:0}
  Step 2:  num=7, complement=19, seen={2:0}     → not found, store {7:1}
  Step 3:  num=11, complement=15, seen={2:0,7:1} → not found, store {11:2}
  Step 4:  num=15, complement=11, seen={2:0,7:1,11:2} → found 11 at idx 2 → return [2, 3]

Why it works:
  - For each num, we check if (target - num) was already seen
  - Map gives O(1) lookup → total O(n) time
  - Works on unsorted arrays (vs Two Pointers which needs sorted input)
*/
