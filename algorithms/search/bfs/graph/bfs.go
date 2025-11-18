// Breadth-First Search (BFS) - Graph Search
// Time: O(V + E) - where V is number of vertices and E is number of edges
// Space: O(V) - for visited map and queue (in worst case, all nodes in queue)

package main

import "fmt"

func FindNodeBFS(graph map[int][]int, start int, target int) (int, bool) {
	// Initialize visited map and queue with starting node
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]  // FIFO
		queue = queue[1:] // dequeue visited neighbor

		if node == target {
			return node, true // we found the target
		}
		// Explore all neighbors
		neighbors := graph[node]
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor) // enqueue unvisited neighbor
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

	node, found := FindNodeBFS(graph, 1, 6)
	if found {
		fmt.Printf("Node %d found\n", node)
	} else {
		fmt.Println("Node not found")
	}
}

/*
How BFS graph search works:

Explores graph level by level using a queue (FIFO - First In, First Out).
Visits all neighbors at current distance before moving to next level.
Uses visited map to prevent cycles and repeated visits.

Parameters:
  - graph: adjacency list representation as map[int][]int
  - start: starting node to begin search
  - target: node value to search for

Returns:
  - int: the found node value, or 0 if not found
  - bool: true if node was found, false otherwise

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

BFS Traversal example (start=1, target=6):

Step 1: queue=[1], visited={1}
        Process node 1 → add neighbors 2,3

Step 2: queue=[2,3], visited={1,2,3}
        Process node 2 → add neighbors 4,5 (1 already visited)

Step 3: queue=[3,4,5], visited={1,2,3,4,5}
        Process node 3 → add neighbor 6 (1 already visited)

Step 4: queue=[4,5,6], visited={1,2,3,4,5,6}
        Process node 4 → no new neighbors

Step 5: queue=[5,6], visited={1,2,3,4,5,6}
        Process node 5 → no new neighbors (2,6 already visited)

Step 6: queue=[6], visited={1,2,3,4,5,6}
        Process node 6 → target found! ✓

Key differences from file system BFS:
  - Works on general graphs (can have cycles)
  - Requires visited map to prevent infinite loops
  - Graph nodes can have multiple incoming edges (multiple "parents")
  - map[int][]int structure instead of os.ReadDir()
*/
