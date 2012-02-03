// Copyright (c) 2011 Graham Anderson <graham@andtech.eu>.
//
// All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pqueue

// pqueue is (for the moment) a very simple priority queue that implements the
// container/heap and sort Interfaces from the Go standard library.

import "container/heap"

type Error string

func (e Error) Error() string {
	return string(e)
}

var (
	QERR_OUT_OF_BOUNDS = Error("This index is out of bounds.")
	QERR_NOT_QUEUEABLE = Error("Object does not implemnt Queueable interface.")
)

// Queue is the basic priority queue structure
type Queue struct {
	heaped     bool
	collection []Queueable
}

// Queueable describes the interface objects must satisfy to be placed in the Q
type Queueable interface {
	Priority() int
}

// NeqQueue returns new Q object with default capacity of 100
func NewQueue() *Queue {
	return &Queue{false, make([]Queueable, 0, 100)}
}

// Len satisfies the sort interface
func (q *Queue) Len() int {
	return len(q.collection)
}

// Less satisfies the sort interface
func (q *Queue) Less(i, j int) bool {
	return q.collection[i].Priority() < q.collection[j].Priority()
}

// Swap satisfies the sort interface
func (q *Queue) Swap(i, j int) {
	q.collection[i], q.collection[j] = q.collection[j], q.collection[i]
}

// Push satisfies the heap interface. NOTE: This method simply appends to the Q
// and does not min-heapify the elements, Consider using Add(x Queueable) or
// container/heap instead e.g. heap.Push(q, x interface{}). Additionally this
// method panics if your Q element does not satisfy the Queueable interface.
// You should protect direct calls to either this or heap.Push() with a
// defer/recover pairing.
func (q *Queue) Push(x interface{}) {
	y, ok := x.(Queueable)
	if !ok {
		panic(QERR_NOT_QUEUEABLE)
	}
	q.collection = append(q.collection, y)
}

func (q *Queue) heapify() {
	if !q.heaped {
		heap.Init(q)
		q.heaped = true
	}
}

// Add a single object to the queue
func (q *Queue) Add(x Queueable) {
	q.heapify()
	heap.Push(q, x)
}

// Remove the highest priority item from the Q
func (q *Queue) Remove() interface{} {
	q.heapify()
	return heap.Pop(q)
}

// PushSlice pushes a slice of queueable objects onto the Q and then min-heaps
// the elements. complexity is O(n) where n = q.Len().
func (q *Queue) AddSlice(s []Queueable) {
	q.collection = append(q.collection, s...)
	q.heapify()
}

// Pop satisfies the heap interface. NOTE: This method pops elements from the Q
// and does not min-heapify the elements, Consider using container/heap instead
// e.g. x, ok := heap.Pop(q).(Queueable)
func (q *Queue) Pop() interface{} {
	x := q.collection[q.Len()-1]
	q.collection = q.collection[0 : q.Len()-1]
	return x
}

// Member returns the queue member at index i
func (q *Queue) Member(i int) (m Queueable, err error) {
	if i < q.Len() {
		m = q.collection[i]
	} else {
		err = QERR_OUT_OF_BOUNDS
	}
	return
}

// Return the full collection (possibly unsorted & not min-heap initialised)
// use heap.Init(q) or sort.Sort(q) if you need the members sorted or min-heaped
func (q *Queue) Collection() []Queueable {
	return q.collection
}
