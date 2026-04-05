// Singly Linked List - Struct-based Implementation
// Space: O(n) - one node per element
//
// Time Complexity:
//   - Prepend:  O(1) - insert at head
//   - Append:   O(1) - insert at tail (tail pointer maintained)
//   - InsertAt: O(n) - traverse to position
//   - Remove:   O(n) - search + unlink
//   - Search:   O(n) - linear scan
//   - Reverse:  O(n) - single pass, pointer reversal
//   - Get:      O(n) - traverse to index
//   - Size:     O(1) - counter maintained
//
// Each node holds a value and a pointer to the next node.
// The list maintains head and tail pointers for O(1) insert at both ends.
// Note: classic singly linked list has only a head pointer (Append is O(n)).
// Tail pointer is a practical optimization; in CS, O(1) insert at both ends
// is typically a property of a doubly linked list or deque.

package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Prepend inserts a value at the beginning of the list.
func (l *LinkedList) Prepend(value int) {
	node := &Node{Value: value, Next: l.Head}
	l.Head = node
	if l.Tail == nil {
		l.Tail = node
	}
	l.size++
}

// Append inserts a value at the end of the list.
func (l *LinkedList) Append(value int) {
	node := &Node{Value: value}
	if l.Tail == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
	l.size++
}

// InsertAt inserts a value at the given index (0-based).
// Returns false if index is out of range.
func (l *LinkedList) InsertAt(index int, value int) bool {
	if index < 0 || index > l.size {
		return false
	}
	if index == 0 {
		l.Prepend(value)
		return true
	}
	if index == l.size {
		l.Append(value)
		return true
	}

	prev := l.Head
	for i := 0; i < index-1; i++ {
		prev = prev.Next
	}
	node := &Node{Value: value, Next: prev.Next}
	prev.Next = node
	l.size++
	return true
}

// Remove deletes the first occurrence of a value.
// Returns true if the value was found and removed.
func (l *LinkedList) Remove(value int) bool {
	if l.Head == nil {
		return false
	}

	// Remove head
	if l.Head.Value == value {
		l.Head = l.Head.Next
		if l.Head == nil {
			l.Tail = nil
		}
		l.size--
		return true
	}

	// Search and unlink
	prev := l.Head
	for prev.Next != nil {
		if prev.Next.Value == value {
			if prev.Next == l.Tail {
				l.Tail = prev
			}
			prev.Next = prev.Next.Next
			l.size--
			return true
		}
		prev = prev.Next
	}
	return false
}

// Search returns true if the value exists in the list.
func (l *LinkedList) Search(value int) bool {
	node := l.Head
	for node != nil {
		if node.Value == value {
			return true
		}
		node = node.Next
	}
	return false
}

// Get returns the value at the given index and true, or 0 and false if out of range.
func (l *LinkedList) Get(index int) (int, bool) {
	if index < 0 || index >= l.size {
		return 0, false
	}
	node := l.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node.Value, true
}

// Reverse reverses the list in place.
func (l *LinkedList) Reverse() {
	l.Tail = l.Head
	var prev *Node
	curr := l.Head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}

// Size returns the number of elements in the list.
func (l *LinkedList) Size() int {
	return l.size
}

// ToSlice returns all values as a slice (head to tail).
func (l *LinkedList) ToSlice() []int {
	result := make([]int, 0, l.size)
	node := l.Head
	for node != nil {
		result = append(result, node.Value)
		node = node.Next
	}
	return result
}

func main() {
	list := NewLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Prepend(0)
	list.InsertAt(2, 99)
	// [0 1 99 2 3]
	fmt.Println(list.ToSlice())

	list.Remove(99)
	// [0 1 2 3]
	fmt.Println(list.ToSlice())

	list.Reverse()
	// [3 2 1 0]
	fmt.Println(list.ToSlice())
}

/*
Linked List structure after Append(1), Append(2), Append(3), Prepend(0), InsertAt(2, 99):

  Head                         Tail
   |                            |
  [0] -> [1] -> [99] -> [2] -> [3] -> nil

After Remove(99):

  Head                 Tail
   |                    |
  [0] -> [1] -> [2] -> [3] -> nil

After Reverse():

  Head                 Tail
   |                    |
  [3] -> [2] -> [1] -> [0] -> nil

Struct-based approach:
  - Each Node holds Value + pointer to Next
  - Head and Tail pointers give O(1) insert at both ends
  - Single direction: can only traverse forward (compare with doubly linked list)
  - Idiomatic Go: explicit pointers, no hidden allocations
*/
