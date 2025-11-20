// Graph - Adjacency List Implementation
// Space: O(V + E) - where V is vertices and E is edges
//
// Time Complexity:
//   - AddVertex:    O(1) - hash map insertion
//   - AddEdge:      O(1) - amortized append to slice
//   - RemoveEdge:   O(E) - linear search through neighbors
//   - RemoveVertex: O(V + E) - must update all neighbors
//   - GetNeighbors: O(1) - direct map access
//   - HasEdge:      O(degree) - linear search through neighbors
//   - Size:         O(1) - map length
//
// This implementation uses an adjacency list representation where each vertex
// maps to a slice of its neighbors. The graph is undirected, meaning edges
// are bidirectional.

package main

import "fmt"

type Graph struct {
	vertices map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int][]int),
	}
}

// Core graph operations

func (g *Graph) AddVertex(vertex int) {
	if _, exists := g.vertices[vertex]; !exists {
		g.vertices[vertex] = []int{}
	}
}

func (g *Graph) AddEdge(v1, v2 int) {
	g.AddVertex(v1)
	g.AddVertex(v2)
	if !g.HasEdge(v1, v2) {
		g.vertices[v1] = append(g.vertices[v1], v2)
	}
	if !g.HasEdge(v2, v1) {
		g.vertices[v2] = append(g.vertices[v2], v1)
	}
}

func (g *Graph) HasEdge(v1, v2 int) bool {
	neighbors := g.vertices[v1]
	for _, neighbor := range neighbors {
		if neighbor == v2 {
			return true
		}
	}
	return false
}

func (g *Graph) RemoveEdge(v1, v2 int) {
	g.vertices[v1] = removeFromSlice(g.vertices[v1], v2)
	g.vertices[v2] = removeFromSlice(g.vertices[v2], v1)
}

func (g *Graph) RemoveVertex(vertex int) {
	for _, neighbor := range g.vertices[vertex] {
		g.vertices[neighbor] = removeFromSlice(g.vertices[neighbor], vertex)
	}
	delete(g.vertices, vertex)
}

// helper
func removeFromSlice(slice []int, value int) []int {
	result := []int{}
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}

func (g *Graph) GetNeighbors(vertex int) []int {
	return g.vertices[vertex]
}

// Utility methods

func (g *Graph) Size() int {
	return len(g.vertices)
}

func (g *Graph) Print() {
	for vertex, neighbors := range g.vertices {
		fmt.Printf("  %d -> %v\n", vertex, neighbors)
	}
}

func main() {
	// Create a new graph
	graph := NewGraph()

	// Add edges (vertices are added automatically)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 6)
	graph.AddEdge(5, 6)

	fmt.Println("Adjacency List Graph:")
	graph.Print()

	fmt.Printf("\nGraph size: %d vertices\n", graph.Size())
	fmt.Printf("Neighbors of vertex 2: %v\n", graph.GetNeighbors(2))
	fmt.Printf("Edge between 1 and 3 exists: %v\n", graph.HasEdge(1, 3))
	fmt.Printf("Edge between 1 and 5 exists: %v\n\n", graph.HasEdge(1, 5))

	// Remove an edge
	graph.RemoveEdge(5, 6)
	fmt.Println("After removing edge (5, 6):")
	graph.Print()

	// Remove a vertex
	fmt.Println("\nAfter removing vertex 2:")
	graph.RemoveVertex(2)
	graph.Print()
}

/*
Graph Visualization:

Initial graph structure:
         1
        / \
       2   3
      / \   \
     4   5---6

Adjacency List representation:
  1 -> [2, 3]
  2 -> [1, 4, 5]
  3 -> [1, 6]
  4 -> [2]
  5 -> [2, 6]
  6 -> [3, 5]

Key Concepts:
   - Most space-efficient for sparse graphs
   - map[int][]int: vertex -> list of neighbors
   - Each edge stored twice (undirected graph)
*/
