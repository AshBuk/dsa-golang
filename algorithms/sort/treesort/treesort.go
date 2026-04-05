// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// Tree Sort: each value inserted into BST, in-order traversal yields sorted output.
// Time: O(n log n) average, O(n²) worst case (sorted input -> not balanced = degenerate "linked list" tree)
// Space: O(n) - for the tree nodes
//
// Based on the tree sort example from
// "The Go Programming Language" by Alan A. A. Donovan & Brian W. Kernighan
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan
// Original source: gopl.io

package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	// values[:0]: len=0, same underlying array and cap.
	// Rewrites original slice in-place without new allocation.
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice
// in-order traversal (left -> node -> right)
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t // equivalent to: return &tree{value: value}
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	values := []int{4, 10, 3, 5, 1}
	Sort(values)
	fmt.Println(values)
}

/*
How Tree Sort works:

Example: [4, 10, 3, 5, 1]

PHASE 1: BUILD BST (Binary Search Tree)
=========================================

Insert 4 (root):
       4

Insert 10 (10 > 4, go right):
       4
        \
        10

Insert 3 (3 < 4, go left):
       4
      / \
     3  10

Insert 5 (5 > 4, go right; 5 < 10, go left):
       4
      / \
     3  10
        /
       5

Insert 1 (1 < 4, go left; 1 < 3, go left):
       4
      / \
     3  10
    /   /
   1   5

PHASE 2: IN-ORDER TRAVERSAL (left → node → right)
===================================================

Start at root (4):
  → go left to 3
    → go left to 1
      → no left child
      → visit 1 ✓
      → no right child
    → visit 3 ✓
    → no right child
  → visit 4 ✓
  → go right to 10
    → go left to 5
      → no left child
      → visit 5 ✓
      → no right child
    → visit 10 ✓
    → no right child

Result: [1, 3, 4, 5, 10] ✓

The trick with values[:0]:
  values[:0] creates a slice with len=0 but same underlying array.
  As appendValues appends sorted elements, they overwrite the
  original array in-place — no extra allocation needed.

Key insights:
- BST property: left < parent ≤ right
- In-order traversal of BST always yields sorted order
- Average case O(n log n): balanced-ish tree, each insert O(log n)
- Worst case O(n²): already sorted input creates a linked list (degenerate tree)
- Not stable: equal elements may change relative order
- Unlike heapsort, requires O(n) extra space for tree nodes
*/
