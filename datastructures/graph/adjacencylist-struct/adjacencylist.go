// Graph - Adjacency List with Struct Nodes
// Space: O(V + E) - where V is vertices and E is edges
//
// Time Complexity:
//   - AddVertex:    O(1) - map insertion
//   - AddEdge:      O(degree) - duplicate check before append
//   - RemoveEdge:   O(degree) - scan neighbor pointers
//   - RemoveVertex: O(V + E) - must update all neighbors
//   - GetNeighbors: O(degree) - copies neighbor IDs into new slice
//   - HasEdge:      O(degree) - scan neighbor pointers
//
// This implementation stores vertices as *Vertex structs with neighbor pointers.
// Each vertex holds its own data and a slice of pointers to neighbors.
// Contrast with the map[int][]int approach in ../adjacencylist-map/adjacencylist.go.

package main

import "fmt"

type Vertex struct {
	ID        int
	Label     string
	Neighbors []*Vertex
}

type Graph struct {
	vertices map[int]*Vertex
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int]*Vertex),
	}
}

// AddVertex creates a vertex with an optional label.
func (g *Graph) AddVertex(id int, label string) *Vertex {
	if v, exists := g.vertices[id]; exists {
		return v
	}
	v := &Vertex{ID: id, Label: label}
	g.vertices[id] = v
	return v
}

// AddEdge creates an undirected edge between two vertices.
func (g *Graph) AddEdge(id1, id2 int) {
	v1 := g.getOrCreate(id1)
	v2 := g.getOrCreate(id2)

	if !hasNeighbor(v1, id2) {
		v1.Neighbors = append(v1.Neighbors, v2)
	}
	if !hasNeighbor(v2, id1) {
		v2.Neighbors = append(v2.Neighbors, v1)
	}
}

func (g *Graph) getOrCreate(id int) *Vertex {
	if v, exists := g.vertices[id]; exists {
		return v
	}
	return g.AddVertex(id, "")
}

func hasNeighbor(v *Vertex, id int) bool {
	for _, n := range v.Neighbors {
		if n.ID == id {
			return true
		}
	}
	return false
}

// HasEdge checks if an edge exists between two vertices.
func (g *Graph) HasEdge(id1, id2 int) bool {
	v1, exists := g.vertices[id1]
	if !exists {
		return false
	}
	return hasNeighbor(v1, id2)
}

// RemoveEdge removes an undirected edge.
func (g *Graph) RemoveEdge(id1, id2 int) {
	if v1, exists := g.vertices[id1]; exists {
		v1.Neighbors = removeNeighbor(v1.Neighbors, id2)
	}
	if v2, exists := g.vertices[id2]; exists {
		v2.Neighbors = removeNeighbor(v2.Neighbors, id1)
	}
}

func removeNeighbor(neighbors []*Vertex, id int) []*Vertex {
	result := make([]*Vertex, 0, len(neighbors))
	for _, n := range neighbors {
		if n.ID != id {
			result = append(result, n)
		}
	}
	return result
}

// RemoveVertex removes a vertex and all its edges.
func (g *Graph) RemoveVertex(id int) {
	v, exists := g.vertices[id]
	if !exists {
		return
	}
	for _, neighbor := range v.Neighbors {
		neighbor.Neighbors = removeNeighbor(neighbor.Neighbors, id)
	}
	delete(g.vertices, id)
}

// GetNeighbors returns neighbor IDs for a given vertex.
func (g *Graph) GetNeighbors(id int) []int {
	v, exists := g.vertices[id]
	if !exists {
		return nil
	}
	ids := make([]int, len(v.Neighbors))
	for i, n := range v.Neighbors {
		ids[i] = n.ID
	}
	return ids
}

// GetVertex returns the vertex struct by ID.
func (g *Graph) GetVertex(id int) (*Vertex, bool) {
	v, exists := g.vertices[id]
	return v, exists
}

// Size returns the number of vertices.
func (g *Graph) Size() int {
	return len(g.vertices)
}

// BFS performs breadth-first search from the given start vertex.
func (g *Graph) BFS(startID int) []int {
	start, exists := g.vertices[startID]
	if !exists {
		return nil
	}

	var result []int
	visited := make(map[int]bool)
	queue := []*Vertex{start}
	visited[start.ID] = true

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		result = append(result, v.ID)

		for _, neighbor := range v.Neighbors {
			if !visited[neighbor.ID] {
				visited[neighbor.ID] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return result
}

// Print displays the graph.
func (g *Graph) Print() {
	for _, v := range g.vertices {
		neighborIDs := make([]int, len(v.Neighbors))
		for i, n := range v.Neighbors {
			neighborIDs[i] = n.ID
		}
		label := ""
		if v.Label != "" {
			label = fmt.Sprintf(" (%s)", v.Label)
		}
		fmt.Printf("  %d%s -> %v\n", v.ID, label, neighborIDs)
	}
}

func main() {
	graph := NewGraph()
	graph.AddVertex(1, "A")
	graph.AddVertex(2, "B")
	graph.AddVertex(3, "C")
	graph.AddVertex(4, "D")
	graph.AddVertex(5, "E")
	graph.AddVertex(6, "F")
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 6)
	graph.AddEdge(5, 6)
	// BFS from vertex 1: [1 2 3 4 5 6]
	fmt.Println(graph.BFS(1))
}

/*
Graph structure:

     1(A)
    /    \
  2(B)   3(C)
  / \      \
4(D) 5(E)--6(F)

Struct-based approach:
  - Each Vertex is a struct: {ID, Label, Neighbors []*Vertex}
  - Neighbors are direct pointers — no ID lookup needed for traversal
  - Vertices can carry arbitrary data (Label, or add more fields)
  - Natural for OOP-style graph algorithms

Comparison with map-based (../adjacencylist.go):
  ┌─────────────────────┬──────────────────────┬─────────────────────┐
  │                     │ map[int][]int         │ Struct *Vertex      │
  ├─────────────────────┼──────────────────────┼─────────────────────┤
  │ Traversal           │ Map lookup per step   │ Direct pointer      │
  │ Node data           │ Separate map needed   │ Fields on struct    │
  │ Memory              │ Compact               │ More allocations    │
  │ Simplicity          │ Minimal boilerplate   │ More code           │
  │ Use case            │ Algorithm problems    │ Rich domain models  │
  └─────────────────────┴──────────────────────┴─────────────────────┘
*/
