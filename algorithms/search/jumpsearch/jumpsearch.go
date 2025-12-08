// Jump Search Algorithm
// Time: O(√n) - jumps (√n) square root of n times -> then linear search within block
// Space: O(1) - only using constant extra space
// Note: requires sorted array

// Use case when "jumping backward (Binary Search) is expensive or impossible, like:
// - Tape storage / sequential files (rewinding backward is expensive)
// - Network streams (data arrives sequentially, can't jump back)
// - Disk I/O where sequential reads are much faster than random seeks

package main

import (
	"fmt"
	"math"
)

func jumpSearch(arr []int, target int) (int, bool) {
	n := len(arr)
	if n == 0 {
		return -1, false
	}
	jumpStep := int(math.Sqrt(float64(n)))
	lo := 0

	for arr[clampToArrLen(jumpStep, n)-1] < target {
		lo = jumpStep
		jumpStep += jumpStep
		if lo >= n {
			return -1, false
		}
	}

	for i := lo; i < clampToArrLen(jumpStep, n); i++ {
		if arr[i] == target {
			return i, true
		}
	}
	return -1, false
}

func clampToArrLen(jumpStep, n int) int {
	if jumpStep < n {
		return jumpStep
	}
	return n // avoid jumping beyond array length
}

func main() {
	fmt.Println(jumpSearch([]int{1, 2, 3, 4, 5, 7, 9, 11, 13}, 11))
}

/*
Example walkthrough: arr = [1, 2, 3, 4, 5, 7, 9, 11, 13], target = 11
n=9, jumpStep=3, lo=0
Jump 1: arr[2]=3 < 11 → lo=3, jumpStep=6
Jump 2: arr[5]=7 < 11 → lo=6, jumpStep=9
Jump 3: arr[8]=13 >= 11 → stop!
Linear search: in block [6..9] → found at index 7
*/
