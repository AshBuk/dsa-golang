// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// Binary Search Tree (BST) - Struct-based Implementation
// Space: O(n) - one node per element
//
// Time Complexity:
//   - Insert:  O(h) - where h is tree height; O(log n) average, O(n) worst (skewed)
//   - Search:  O(h) - follows one path from root to leaf
//   - Remove:  O(h) - find node + restructure
//   - Min/Max: O(h) - traverse leftmost/rightmost path
//   - InOrder: O(n) - visits every node exactly once
//
// BST invariant: for every node, all values in left subtree < node.value < all values in right subtree.
// This implementation uses struct with pointers — the idiomatic Go approach for tree structures.

package main

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

type BST struct {
	Root *Node
}

func NewBST() *BST {
	return &BST{}
}

// Insert adds a value to the BST maintaining the BST invariant.
func (t *BST) Insert(value int) {
	t.Root = insert(t.Root, value)
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value}
	}
	if value < node.Value {
		node.Left = insert(node.Left, value)
	} else if value > node.Value {
		node.Right = insert(node.Right, value)
	}
	// duplicate values are ignored
	return node
}

// Search returns true if value exists in the tree.
func (t *BST) Search(value int) bool {
	return search(t.Root, value)
}

func search(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	}
	if value < node.Value {
		return search(node.Left, value)
	}
	return search(node.Right, value)
}

// Remove deletes a value from the BST.
func (t *BST) Remove(value int) {
	t.Root = remove(t.Root, value)
}

func remove(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	if value < node.Value {
		node.Left = remove(node.Left, value)
	} else if value > node.Value {
		node.Right = remove(node.Right, value)
	} else {
		// Found the node to remove
		// Case 1: leaf node
		if node.Left == nil && node.Right == nil {
			return nil
		}
		// Case 2: one child
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}
		// Case 3: two children — replace with in-order successor (min of right subtree)
		successor := findMin(node.Right)
		node.Value = successor.Value
		node.Right = remove(node.Right, successor.Value)
	}
	return node
}

// Min returns the smallest value in the tree.
func (t *BST) Min() (int, bool) {
	if t.Root == nil {
		return 0, false
	}
	return findMin(t.Root).Value, true
}

func findMin(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// Max returns the largest value in the tree.
func (t *BST) Max() (int, bool) {
	if t.Root == nil {
		return 0, false
	}
	node := t.Root
	for node.Right != nil {
		node = node.Right
	}
	return node.Value, true
}

// InOrder returns all values in sorted order (left → root → right).
func (t *BST) InOrder() []int {
	var result []int
	inOrder(t.Root, &result)
	return result
}

func inOrder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	inOrder(node.Left, result)
	*result = append(*result, node.Value)
	inOrder(node.Right, result)
}

// Height returns the height of the tree (longest path from root to leaf).
func (t *BST) Height() int {
	return height(t.Root)
}

func height(node *Node) int {
	if node == nil {
		return -1
	}
	leftH := height(node.Left)
	rightH := height(node.Right)
	if leftH > rightH {
		return leftH + 1
	}
	return rightH + 1
}

func main() {
	bst := NewBST()
	for _, v := range []int{8, 3, 10, 1, 6, 14, 4, 7, 13} {
		bst.Insert(v)
	}
	// In-order: [1 3 4 6 7 8 10 13 14]
	fmt.Println(bst.InOrder())
}

/*
BST Structure after insertions:

            8
          /   \
         3     10
        / \      \
       1   6     14
          / \    /
         4   7  13

In-order traversal visits nodes in sorted order:
  1, 3, 4, 6, 7, 8, 10, 13, 14

Remove cases:
  ✓ Leaf (no children): simply remove the node
  ✓ One child: replace node with its child
  ✓ Two children: replace with in-order successor (smallest in right subtree)

Struct-based approach:
  - Each Node holds Value + pointers to Left and Right children
  - Recursive functions naturally follow the tree structure
  - Idiomatic for Go: explicit, type-safe, clean recursion
  - Compare with map-based tree (see ../nary-map/) for an alternative representation
*/
