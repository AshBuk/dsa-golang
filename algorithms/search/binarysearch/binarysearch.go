// Binary Search Algorithm
// Time: O(log n) - divides search space in half each iteration
// Space: O(1) - only using constant extra space
// Note: requires sorted array

package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) (int, bool) {
	lo := 0
	hi := len(arr) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2 // prevents int overflow vs (lo+hi)/2;
		// not required in Go due to 64-bit ints, but good practice
		if target == arr[mid] {
			return mid, true
		} else if target < arr[mid] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1, false
}

func main() {
	fmt.Println(binarySearch([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8))
}
