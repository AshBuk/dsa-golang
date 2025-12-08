// Time: O(n) - single pass, map lookup is O(1)
// Space: O(n) - map stores up to n unique elements + result array

package main

import (
	"fmt"
)

func removeDuplicates(arr []int) []int {
	seen := make(map[int]bool)
	uniqueElements := []int{}

	for _, v := range arr {
		if !seen[v] {
			seen[v] = true
			uniqueElements = append(uniqueElements, v)
		}
	}
	return uniqueElements
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2, 3, 3, 3, 4, 5}))
}

// Time: O(n²) - nested loops
// Space: O(n) - worst case: all elements unique, result array size = n

// func removeDuplicates(arr []int) []int {
// 	uniqElements := []int{}
// 	for i := 0; i < len(arr); i++ {
// 		found := false
// 		for j := 0; j < len(uniqElements); j++ {
// 			if uniqElements[j] == arr[i] {
// 				found = true
// 				break
// 			}
// 		}
// 		if !found {
// 			uniqElements = append(uniqElements, arr[i])
// 		}
// 	}
// 	return uniqElements
// }
