// Queue - Slice-based Implementation (FIFO)
// Space: O(n) - one slot per element in the underlying slice
//
// Time Complexity:
//   - Enqueue:  O(1) amortized - append to slice
//   - Dequeue:  O(1) - advance front index
//   - Peek:     O(1) - read front element
//   - Size:     O(1) - counter maintained
//   - IsEmpty:  O(1) - check counter
//
// Slice-based approach for queues is simple, cache-friendly,
// and avoids pointer overhead of a linked list.
// Uses a front index to avoid costly slice shifting on dequeue.

package main

import "fmt"

type Queue struct {
	items []int
	front int
}

func NewQueue() *Queue {
	return &Queue{}
}

// Enqueue adds a value to the back of the queue.
func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

// Dequeue removes and returns the front value.
// Returns the value and true, or 0 and false if the queue is empty.
func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	value := q.items[q.front]
	q.front++

	// Reclaim memory when more than half the slice is unused
	if q.front > len(q.items)/2 {
		q.items = append([]int{}, q.items[q.front:]...)
		q.front = 0
	}
	return value, true
}

// Peek returns the front value without removing it.
func (q *Queue) Peek() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.items[q.front], true
}

// Size returns the number of elements in the queue.
func (q *Queue) Size() int {
	return len(q.items) - q.front
}

// IsEmpty returns true if the queue has no elements.
func (q *Queue) IsEmpty() bool {
	return q.front >= len(q.items)
}

// ToSlice returns all elements from front to back.
func (q *Queue) ToSlice() []int {
	result := make([]int, q.Size())
	copy(result, q.items[q.front:])
	return result
}

func main() {
	q := NewQueue()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	// [10 20 30]
	fmt.Println(q.ToSlice())

	val, _ := q.Dequeue()
	// Dequeued: 10
	fmt.Println("Dequeued:", val)

	val, _ = q.Peek()
	// Front: 20
	fmt.Println("Front:", val)

	q.Enqueue(40)
	// [20 30 40]
	fmt.Println(q.ToSlice())
}

/*
Queue operations (FIFO -- First In, First Out):

  Enqueue(10), Enqueue(20), Enqueue(30):

    front       back
      |          |
    [ 10 | 20 | 30 ]

  Dequeue() -> 10:

         front  back
           |     |
    [    | 20 | 30 ]

  Enqueue(40):

         front       back
           |          |
    [    | 20 | 30 | 40 ]

Slice-based approach:
  - Uses []int with a front index — no pointer chasing
  - Enqueue = append, Dequeue = advance index
  - Periodic compaction reclaims unused space at the front
  - Cache-friendly: contiguous memory layout
  - Compare with pointer-based (linked list) for O(1) dequeue without compaction
*/
