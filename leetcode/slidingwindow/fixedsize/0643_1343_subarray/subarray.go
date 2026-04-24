// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// LeetCode 643: Maximum Average Subarray I
// https://leetcode.com/problems/maximum-average-subarray-i/
//
// You are given an integer array nums consisting of n elements, and an integer k.
// Find a contiguous subarray whose length is equal to k that has the maximum average
// value and return this value.
//
// Time: O(n) - single pass after initial window sum
// Space: O(1) - only scalar variables used
//
// LeetCode 1343: Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold
// https://leetcode.com/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/
//
// Given an array of integers arr and two integers k and threshold, return the number
// of sub-arrays of size k and average greater than or equal to threshold.
//
// Time: O(n) - single pass after initial window sum
// Space: O(1) - only scalar variables used

package main

import "fmt"

// 643: maximum average over a fixed window of size k
func findMaxAverage(nums []int, k int) float64 {
	var windowSum int
	for i := range k {
		windowSum += nums[i]
	}
	maxSum := windowSum
	for right := k; right < len(nums); right++ {
		left := right - k
		windowSum += nums[right] - nums[left]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return float64(maxSum) / float64(k)
}

// 1343: count fixed-size windows whose average >= threshold
// Trick: compare windowSum >= threshold * k to stay in integer arithmetic.
func numOfSubarrays(arr []int, k int, threshold int) int {
	var windowSum int
	for i := range k {
		windowSum += arr[i]
	}
	var count int
	target := threshold * k
	if windowSum >= target {
		count++
	}
	for right := k; right < len(arr); right++ {
		left := right - k
		windowSum += arr[right] - arr[left]
		if windowSum >= target {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)) // 12.75
	fmt.Println(findMaxAverage([]int{5}, 1))                    // 5

	fmt.Println(numOfSubarrays([]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4))             // 3
	fmt.Println(numOfSubarrays([]int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, 3, 5)) // 6
}

/*
-#643 — Fixed-size sliding window: nums = [1, 12, -5, -6, 50, 3], k = 4

  Init window [0..3]:  windowSum = 1+12+(-5)+(-6) = 2,  maxSum = 2

  right=4, left=0:  windowSum = 2 - nums[0] + nums[4] = 2 - 1 + 50 = 51   → maxSum = 51
  right=5, left=1:  windowSum = 51 - nums[1] + nums[5] = 51 - 12 + 3 = 42 → maxSum = 51

  return 51 / 4 = 12.75

Why it works:
  - Instead of recomputing the sum from scratch each step (O(n*k)),
    we slide by subtracting the element leaving the window and adding the one entering
  - Divide only at the end — comparing sums is equivalent to comparing averages

-#1343 — Fixed-size sliding window: arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4

  target = threshold * k = 12   (compare sums, stay in integer arithmetic)

  Init window [0..2]:  windowSum = 2+2+2 = 6     →  6 >= 12? no    count = 0

  right=3, left=0:  windowSum = 6 - arr[0] + arr[3] = 6 - 2 + 2 = 6      →  no     count = 0
  right=4, left=1:  windowSum = 6 - arr[1] + arr[4] = 6 - 2 + 5 = 9      →  no     count = 0
  right=5, left=2:  windowSum = 9 - arr[2] + arr[5] = 9 - 2 + 5 = 12     →  yes    count = 1
  right=6, left=3:  windowSum = 12 - arr[3] + arr[6] = 12 - 2 + 5 = 15   →  yes    count = 2
  right=7, left=4:  windowSum = 15 - arr[4] + arr[7] = 15 - 5 + 8 = 18   →  yes    count = 3

  return 3

Key difference between #643 and #1343:
  -#643: track the maximum windowSum, then divide by k at the end
  -#1343: count windows where windowSum >= threshold * k (integer-only comparison)
  - Identical sliding mechanics: windowSum = windowSum - arr[left] + arr[right]
*/
