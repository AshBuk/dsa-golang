// Jump Search Algorithm
// Time: O(√n) - jumps √n blocks of size √n, then linear search within block
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
	step := int(math.Sqrt(float64(n)))
	prev := 0
	// Jump forward
	for arr[min(step, n)-1] < target {
		prev = step
		step += int(math.Sqrt(float64(n))) // fixed step size
		if prev >= n {
			return -1, false
		}
	}
	// Linear search within the block [prev .. min(step, n)]
	for i := prev; i < min(step, n); i++ {
		if arr[i] == target {
			return i, true
		}
	}
	return -1, false
}

// Note: custom min is used for Go < 1.21 compatibility
// and to make algorithm boundaries explicit.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(jumpSearch([]int{1, 2, 3, 4, 5, 7, 9, 11, 13}, 11))
}

/*
Example walkthrough: arr = [1, 2, 3, 4, 5, 7, 9, 11, 13], target = 11
n=9, step=3 (√9=3), prev=0
Jump 1: arr[2]=3 < 11 → prev=3, step=6
Jump 2: arr[5]=7 < 11 → prev=6, step=9
Jump 3: arr[8]=13 >= 11 → stop!
Linear search: in block [6..9] → found at index 7
*/
