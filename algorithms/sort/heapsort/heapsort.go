// Heap Sort Algorithm
// Time: O(n log n) - build heap O(n) + n extractions O(log n) each
// Space: O(1) - sorts in-place
// Note: uses max heap to sort in ascending order

package main

import "fmt"

func heapSort(arr []int) []int {
	n := len(arr)

	// Build max heap from array
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from heap one by one
	for i := n - 1; i > 0; i-- {
		// Move current root (max) to end
		arr[0], arr[i] = arr[i], arr[0]
		// Heapify reduced heap
		heapify(arr, i, 0)
	}

	return arr
}

// heapify maintains max heap property for subtree rooted at index i
func heapify(arr []int, heapSize int, rootIdx int) {
	largest := rootIdx
	left := 2*rootIdx + 1
	right := 2*rootIdx + 2

	// If left child is larger than root
	if left < heapSize && arr[left] > arr[largest] {
		largest = left
	}

	// If right child is larger than largest so far
	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}

	// If largest is not root, swap and continue heapifying
	if largest != rootIdx {
		arr[rootIdx], arr[largest] = arr[largest], arr[rootIdx]
		heapify(arr, heapSize, largest)
	}
}

func main() {
	fmt.Println(heapSort([]int{4, 10, 3, 5, 1}))
}

/*
How Heap Sort works:

Example: [4, 10, 3, 5, 1]

PHASE 1: BUILD MAX HEAP
========================

Initial array visualized as binary tree:
       4
      / \
    10   3
   /  \
  5    1

Indices: parent = i, left child = 2*i+1, right child = 2*i+2

Step 1: Heapify from i = n/2-1 = 1 (node with value 10)
       4
      / \
   →10   3    ← check if 10 is larger than its children (5, 1)
   /  \          10 > 5 and 10 > 1, so no swap needed
  5    1

Step 2: Heapify from i = 0 (root node with value 4)
     →4         ← check if 4 is larger than its children (10, 3)
      / \          4 < 10, so swap them
    10   3
   /  \
  5    1

After swap:
      10         ← 4 moved down
      / \
   →4   3       ← continue heapifying at index 1
   /  \            check if 4 is larger than its children (5, 1)
  5    1           4 < 5, so swap them

Final max heap:
      10         ← largest element at root
      / \
     5   3
    /  \
   4    1

Array representation: [10, 5, 3, 4, 1]

PHASE 2: EXTRACT AND SORT
==========================

Step 1: Swap root (10) with last element (1)
Array: [1, 5, 3, 4, | 10]  ← 10 is now in final position
Heap size reduced to 4, heapify from root:

       1         ← swap with largest child (5)
      / \
     5   3
    /
   4

After heapify:
       5
      / \
     4   3
    /
   1
Array: [5, 4, 3, 1, | 10]

Step 2: Swap root (5) with last unsorted element (1)
Array: [1, 4, 3, | 5, 10]  ← 5 is in final position
Heapify:

       1         ← swap with largest child (4)
      / \
     4   3

After heapify:
       4
      / \
     1   3
Array: [4, 1, 3, | 5, 10]

Step 3: Swap root (4) with last unsorted element (3)
Array: [3, 1, | 4, 5, 10]
Heapify:

       3         ← 3 > 1, no swap needed
      /
     1

Array: [3, 1, | 4, 5, 10]

Step 4: Swap root (3) with last unsorted element (1)
Array: [1, | 3, 4, 5, 10]  ← only one element left, done!

FINAL SORTED ARRAY: [1, 3, 4, 5, 10] ✓

Key insights:
- Max heap: parent ≥ children (root contains maximum)
- Build heap: O(n) - start from last parent node, heapify upwards
- Extract max: O(log n) per extraction, n extractions = O(n log n)
- In-place: no extra arrays needed, only O(1) extra space
- Not stable: equal elements may change relative order
*/
