// Linear Search Algorithm
// Time: O(n) - worst case we need to check every element
// Space: O(1) - only using constant extra space

package main

import "fmt"

func linearSearch(arr []int, target int) (int, bool) {
	for idx, v := range arr {
		if v == target {
			return idx, true
		}
	}
	return -1, false
}

func main() {
	fmt.Println(linearSearch([]int{1, 4, 3, 6, 2, 7}, -1))
}
