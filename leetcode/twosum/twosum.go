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
	numToIdx := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		// If complement exists in map, we found the pair
		if idx, ok := numToIdx[complement]; ok {
			return []int{idx, i}
		}
		// Store current number and its index for future lookups
		numToIdx[num] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 26))
}

// Brute Force
// Time: O(n²)
// Space: O(1)
// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return nil
// }
