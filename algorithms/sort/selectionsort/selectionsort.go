// Selection Sort Algorithm
// Time Complexity: O(n²) - nested loops iterate through the array
// Space Complexity: O(1) - only uses constant extra space for variables

package main

import "fmt"

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i { // swap only if min. element isn't already in correct position
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
	return arr
}

func main() {
	fmt.Println(selectionSort([]int{1, 6, 5, 4, 7, 2, 3, 8}))
}

/*
How selection sort works:

Example: [1, 6, 5, 4, 7, 2, 3, 8]

ITERATION 1 (i=0):
    [1, 6, 5, 4, 7, 2, 3, 8]
     ↑ current (i=0)
    minIdx = i = 0
    Search for min: j=1..7, compare arr[j] with arr[minIdx]
    minIdx = 0 (arr[0]=1 is already minimum)
    [1, 6, 5, 4, 7, 2, 3, 8]  ← already in correct position

ITERATION 2 (i=1):
    [1, 6, 5, 4, 7, 2, 3, 8]
        ↑ current (i=1)
    minIdx = i = 1
    Search for min: j=2..7
    j=2: arr[2]=5 < arr[minIdx]=6 → minIdx = 2
    j=3: arr[3]=4 < arr[minIdx]=5 → minIdx = 3
    j=4: arr[4]=7 > arr[minIdx]=4 ✗
    j=5: arr[5]=2 < arr[minIdx]=4 → minIdx = 5
    j=6: arr[6]=3 > arr[minIdx]=2 ✗
    j=7: arr[7]=8 > arr[minIdx]=2 ✗
    minIdx = 5 (arr[5]=2 is minimum)
    Swap arr[i] with arr[minIdx]: arr[1] ↔ arr[5]
    [1, 2, 5, 4, 7, 6, 3, 8]

	and so on...

Key points:
- For each position i, finds minimum element in unsorted portion (j = i+1..n-1)
- Swaps minimum element with element at position i
- Builds sorted portion from left to right
- Always O(n²) - no early exit optimization
- Simple but inefficient for large arrays
*/
