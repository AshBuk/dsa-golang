// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// LeetCode 88: Merge Sorted Array
// https://leetcode.com/problems/merge-sorted-array/
//
// merge combines two sorted arrays nums1 and nums2 in-place into nums1.
// The function uses a two-pointer approach starting from the end of both arrays,
// comparing elements and placing them in the correct position from right to left.
// This avoids the need for extra space and prevents overwriting elements in nums1.
//
// Parameters:
//   - nums1: destination array with length m+n, where first m elements are valid
//   - m: number of valid elements in nums1
//   - nums2: source array with n elements
//   - n: number of elements in nums2
//
// Time complexity: O(m+n), Space complexity: O(1)

package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1    // pointer for last valid element in nums1
	p2 := n - 1    // pointer for last element in nums2
	p := m + n - 1 // pointer for last position in nums1

	// Merge from right to left, comparing elements from both arrays
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
	// Copy remaining elements from nums2 if any
	// (no need to copy from nums1 as they're already in place)
	if p2 >= 0 {
		copy(nums1[:p2+1], nums2[:p2+1])
	}
}

func main() {
	nums1 := []int{1, 3, 3, 5, 5, 9, 0, 0, 0, 0, 0}
	nums2 := []int{0, 2, 3, 4, 66}

	fmt.Println("Before merge:", nums1, nums2)
	merge(nums1, 6, nums2, 5)
	fmt.Println("After merge:", nums1)
}
