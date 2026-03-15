// N-ary Tree - Map-based Implementation
// Space: O(n) - one entry per node in the map
//
// Time Complexity:
//   - AddNode:    O(1) - map insertion + append to parent's children
//   - GetChildren: O(1) - direct map access
//   - BFS:        O(n) - visits every node
//   - DFS:        O(n) - visits every node
//   - FindPath:   O(n) - worst case traverses entire tree
//   - Height:     O(n) - visits every node
//
// Map-based representation: tree stored as map[int][]int where key is node ID
// and value is a slice of child IDs. Useful when nodes are identified by index
// or when the tree is built from edge lists (e.g., graph problems).

package main

import "fmt"

type Tree struct {
	children map[int][]int
	values   map[int]string // optional payload per node
	root     int
}

func NewTree(root int) *Tree {
	t := &Tree{
		children: make(map[int][]int),
		values:   make(map[int]string),
		root:     root,
	}
	t.children[root] = []int{}
	return t
}

// AddNode adds a child node under the given parent.
func (t *Tree) AddNode(parent, child int) {
	t.children[parent] = append(t.children[parent], child)
	if _, exists := t.children[child]; !exists {
		t.children[child] = []int{}
	}
}

// SetValue assigns an optional label/value to a node.
func (t *Tree) SetValue(node int, value string) {
	t.values[node] = value
}

// GetChildren returns the children of a node.
func (t *Tree) GetChildren(node int) []int {
	return t.children[node]
}

// Size returns the total number of nodes.
func (t *Tree) Size() int {
	return len(t.children)
}

// BFS returns nodes in breadth-first order.
func (t *Tree) BFS() []int {
	var result []int
	queue := []int{t.root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)
		queue = append(queue, t.children[node]...)
	}
	return result
}

// DFS returns nodes in depth-first (pre-order) order.
func (t *Tree) DFS() []int {
	var result []int
	t.dfs(t.root, &result)
	return result
}

func (t *Tree) dfs(node int, result *[]int) {
	*result = append(*result, node)
	for _, child := range t.children[node] {
		t.dfs(child, result)
	}
}

// FindPath returns the path from root to target, or nil if not found.
func (t *Tree) FindPath(target int) []int {
	var path []int
	if t.findPath(t.root, target, &path) {
		return path
	}
	return nil
}

func (t *Tree) findPath(node, target int, path *[]int) bool {
	*path = append(*path, node)
	if node == target {
		return true
	}
	for _, child := range t.children[node] {
		if t.findPath(child, target, path) {
			return true
		}
	}
	*path = (*path)[:len(*path)-1]
	return false
}

// Height returns the height of the tree.
func (t *Tree) Height() int {
	return t.height(t.root)
}

func (t *Tree) height(node int) int {
	children := t.children[node]
	if len(children) == 0 {
		return 0
	}
	maxH := 0
	for _, child := range children {
		h := t.height(child)
		if h > maxH {
			maxH = h
		}
	}
	return maxH + 1
}

// Print displays the tree structure with indentation.
func (t *Tree) Print() {
	t.printNode(t.root, "", true)
}

func (t *Tree) printNode(node int, prefix string, isLast bool) {
	connector := "├── "
	if isLast {
		connector = "└── "
	}
	if node == t.root {
		connector = ""
	}

	label := fmt.Sprintf("%d", node)
	if v, ok := t.values[node]; ok {
		label = fmt.Sprintf("%d (%s)", node, v)
	}
	fmt.Printf("%s%s%s\n", prefix, connector, label)

	childPrefix := prefix
	if node != t.root {
		if isLast {
			childPrefix += "    "
		} else {
			childPrefix += "│   "
		}
	}

	children := t.children[node]
	for i, child := range children {
		t.printNode(child, childPrefix, i == len(children)-1)
	}
}

func main() {
	tree := NewTree(0)
	tree.AddNode(0, 1)
	tree.AddNode(0, 2)
	tree.AddNode(0, 3)
	tree.AddNode(1, 4)
	tree.AddNode(1, 5)
	tree.AddNode(1, 6)
	tree.AddNode(6, 7)
	tree.AddNode(6, 8)
	tree.AddNode(2, 9)
	// BFS: [0 1 2 3 4 5 6 9 7 8]
	fmt.Println(tree.BFS())
}

/*
Tree structure (N-ary, map-based):

  0 (root)
  ├── 1 (src)
  │   ├── 4 (main.go)
  │   ├── 5 (utils.go)
  │   └── 6 (handlers)
  │       ├── 7 (auth.go)
  │       └── 8 (api.go)
  ├── 2 (docs)
  │   └── 9 (README.md)
  └── 3 (tests)

Map representation:
  children: {0: [1,2,3], 1: [4,5,6], 2: [9], 3: [], 6: [7,8], ...}
  values:   {0: "root", 1: "src", 4: "main.go", ...}

Map-based approach:
  - Nodes identified by int IDs, relationships stored in map[int][]int
  - Easy to build from edge lists or adjacency data
  - Natural for BFS (queue-based traversal)
  - Flexible: any number of children per node (N-ary)
  - Compare with struct-based tree (see ../bst/) for pointer-based approach
*/
