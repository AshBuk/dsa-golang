// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// Depth-First Search (DFS) - Graph Search (Iterative with Stack)
// Time: O(V + E) - where V is number of vertices and E is number of edges
// Space: O(V) - for visited map and explicit stack

package main

import "fmt"

func FindNodeDFSIterative(graph map[int][]int, start int, target int) (int, bool) {
	visited := make(map[int]bool)
	stack := []int{start} // ← STACK (LIFO)

	for len(stack) > 0 {
		node := stack[len(stack)-1]  // get last node
		stack = stack[:len(stack)-1] // and pop from the stack

		if visited[node] { // skip if already visited
			continue
		}
		visited[node] = true // mark current node as visited
		if node == target {
			return node, true // we found the target
		}
		// Explore all neighbors
		neighbors := graph[node]
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				stack = append(stack, neighbor) // push unvisited neighbor to stack
			}
		}
	}
	return 0, false
}

func main() {
	// Example graph as adjacency list
	graph := map[int][]int{
		1: {2, 3},
		2: {1, 4, 5},
		3: {1, 6},
		4: {2},
		5: {2, 6},
		6: {3, 5},
	}

	node, found := FindNodeDFSIterative(graph, 1, 6)
	if found {
		fmt.Printf("Node %d found\n", node)
	} else {
		fmt.Println("Node not found")
	}
}

/*
How DFS iterative graph search works:

Uses an explicit stack (slice) instead of recursion to achieve depth-first traversal.
Explores each branch fully before backtracking (depth-first traversal).
Uses visited map to prevent cycles and repeated visits.

Parameters:
  - graph: adjacency list representation as map[int][]int
  - start: starting node to begin search
  - target: node value to search for

Returns:
  - int: the found node value, or 0 if not found
  - bool: true if node was found, false otherwise

Key difference from recursive version:
  - Uses explicit stack ([]int) instead of call stack
  - No risk of stack overflow on deep graphs
  - More control over memory usage
  - Slightly more verbose code

Stack operations:
  - Push: stack = append(stack, item)
  - Pop: item := stack[len(stack)-1]; stack = stack[:len(stack)-1]
  - LIFO (Last In, First Out) - takes most recently added item

Graph representation - Adjacency List:

map[int][]int is a map where:
  - key (int): node identifier
  - value ([]int): slice of neighbor nodes that this node connects to

Example graph from main():

    graph := map[int][]int{
        1: {2, 3},
        2: {1, 4, 5},
        3: {1, 6},
        4: {2},
        5: {2, 6},
        6: {3, 5},
    }

Visual representation:

         1
        / \
       2   3
      /|   |\
     4 5   6 |
        \ /  |
         *---*

Adjacency list breakdown:

    Node 1 → neighbors [2, 3]       (1 connects to 2 and 3)
    Node 2 → neighbors [1, 4, 5]    (2 connects to 1, 4, and 5)
    Node 3 → neighbors [1, 6]       (3 connects to 1 and 6)
    Node 4 → neighbors [2]          (4 connects to 2)
    Node 5 → neighbors [2, 6]       (5 connects to 2 and 6)
    Node 6 → neighbors [3, 5]       (6 connects to 3 and 5)

DFS Iterative Traversal example (start=1, target=6):

Step 1: stack=[1], visited={}
        Pop node 1 → mark visited → add neighbors 2,3

Step 2: stack=[2,3], visited={1}
        Pop node 3 (LIFO!) → mark visited → add neighbor 6 (1 already visited)

Step 3: stack=[2,6], visited={1,3}
        Pop node 6 → mark visited → target found! ✓

Note: Order may differ from recursive version due to stack operations,
but both achieve depth-first traversal and find the target.
*/
