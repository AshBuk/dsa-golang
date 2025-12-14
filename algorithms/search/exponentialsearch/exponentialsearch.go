// Exponential Search Algorithm
// Time: O(log n) - exponentially expands range, then binary search within block
// Space: O(1) - only using constant extra space
// Note: requires sorted array

// Use case:
// - Unbounded/infinite arrays where size is unknown
// - When target is likely near the beginning (faster than pure binary search)
// - External data sources where sequential forward access is cheap

package main

import "fmt"

func exponentialSearch(arr []int, target int) (int, bool) {
	n := len(arr)
	if n == 0 {
		return -1, false
	}
	// Edge case: check first element
	if arr[0] == target {
		return 0, true
	}
	// Exponential expansion
	bound := 1
	for bound < n && arr[bound] < target {
		bound *= 2
	}
	// Binary search within range
	lo := bound / 2
	hi := min(bound, n-1)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if arr[mid] == target {
			return mid, true
		} else if arr[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1, false
}

func main() {
	fmt.Println(exponentialSearch([]int{1, 2, 3, 4, 5, 7, 9, 11, 13}, 11))
}

/*
Example walkthrough: arr = [1, 2, 3, 4, 5, 7, 9, 11, 13], target = 11
n=9

Phase 1 - Exponential expansion:
bound=1: arr[1]=2 < 11 → bound=2
bound=2: arr[2]=3 < 11 → bound=4
bound=4: arr[4]=5 < 11 → bound=8
bound=8: arr[8]=13 >= 11 → stop!

Phase 2 - Binary search in range [4, 8]:
lo=4, hi=8, mid=6: arr[6]=9 < 11 → lo=7
lo=7, hi=8, mid=7: arr[7]=11 == 11 → found at index 7!
*/
