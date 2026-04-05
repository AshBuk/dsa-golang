// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// Graph - Adjacency List Implementation
// Space: O(V + E) - where V is vertices and E is edges
//
// Time Complexity:
//   - AddVertex:    O(1) - hash map insertion
//   - AddEdge:      O(degree) - duplicate check before append
//   - RemoveEdge:   O(degree) - linear search through neighbors
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
	graph := NewGraph()
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 6)
	graph.AddEdge(5, 6)
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
