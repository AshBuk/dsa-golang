// Depth-First Search (DFS) - Graph Search
// Time: O(V + E) - where V is number of vertices and E is number of edges
// Space: O(V) - for visited map and call stack (recursion depth)

package main

import "fmt"

func FindNodeDFS(graph map[int][]int, start int, target int) (int, bool) {
	visited := make(map[int]bool)
	return findNodeDFSRecursive(graph, start, target, visited)
}

func findNodeDFSRecursive(graph map[int][]int, node int, target int, visited map[int]bool) (int, bool) {
	visited[node] = true // mark current node as visited

	if node == target {
		return node, true // we found the target
	}
	// Explore all neighbors
	neighbors := graph[node]
	for _, neighbor := range neighbors {
		if !visited[neighbor] {
			if foundNode, found := findNodeDFSRecursive(graph, neighbor, target, visited); found {
				return foundNode, true
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

	node, found := FindNodeDFS(graph, 1, 6)
	if found {
		fmt.Printf("Node %d found\n", node)
	} else {
		fmt.Println("Node not found")
	}
}

/*
How DFS graph search works:

Explores each branch fully before backtracking (depth-first traversal).
Uses recursion to visit neighbors immediately when encountered.
Uses visited map to prevent cycles and repeated visits.

Parameters:
  - graph: adjacency list representation (map of node -> neighbors)
  - start: starting node to begin search
  - target: node value to search for

Returns:
  - int: the found node value, or 0 if not found
  - bool: true if node was found, false otherwise

Graph representation - Adjacency List:

map[int][]int where key is node ID, value is slice of neighbors.

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

Adjacency breakdown:
    1 → [2, 3]
    2 → [1, 4, 5]
    3 → [1, 6]
    4 → [2]
    5 → [2, 6]
    6 → [3, 5]

DFS Traversal (start=1, target=6):
    Visit 1 → Recurse 2 → Recurse 4 → Backtrack → Recurse 5 → Recurse 6 ✓
    Path: 1 → 2 → 4 (backtrack) → 5 → 6 (found!)

Key points:
  - Works on graphs (can have cycles)
  - Requires visited map to prevent infinite recursion
  - Goes deep before exploring siblings
*/
