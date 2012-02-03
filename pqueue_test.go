// Copyright (c) 2011 Graham Anderson <graham@andtech.eu>.
//
// All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pqueue_test

import (
	"container/heap"
	"fmt"
	pq "github.com/gnanderson/pqueue"
	"math/rand"
	"sort"
	"testing"
)

type message struct {
	importance int
}

func (m *message) Priority() int {
	return m.importance
}

func initCollection() *pq.Queue {
	q := pq.NewQueue()
	for i := 0; i < 100; i++ {
		q.Push(&message{rand.Int()})
	}
	return q
}

func initSmallCollection(n int) *pq.Queue {
	q := pq.NewQueue()
	for i := 0; i < 10; i++ {
		q.Push(&message{rand.Intn(n)})
	}
	return q
}

func TestAddElement(t *testing.T) {
	q := initCollection()
	fmt.Println("Length:", q.Len())
	if q.Len() != 100 {
		t.Fatalf("Add() fails, wrong number of elements.")
	}
}

func TestLess(t *testing.T) {
	q := pq.NewQueue()
	q.Push(&message{132})
	q.Push(&message{256})
	if !q.Less(0, 1) {
		t.Fatalf("Less() fails.")
	}
}

func TestSwap(t *testing.T) {
	q := pq.NewQueue()
	q.Push(&message{132})
	q.Push(&message{256})
	q.Swap(0, 1)
	x, _ := q.Member(0)
	y, _ := q.Member(1)
	if x.Priority() != 256 || y.Priority() != 132 {
		t.Fatalf("Swap() fails.")
	}
}

func TestSort(t *testing.T) {
	q := initCollection()
	sort.Sort(q)

	for i := 1; i < q.Len(); i++ {
		x, err := q.Member(i - 1)
		if err != nil {
			t.Fatalf("Invalid number of elements")
		}
		y, err := q.Member(i)
		if err != nil {
			t.Fatalf("Invalid number of elements")
		}
		if x.Priority() > y.Priority() {
			t.Fatalf("Sort fails, not sorted")
		}
	}
}

func TestPushSlice(t *testing.T) {
	q := initCollection()
	tmp := make([]pq.Queueable, 10)
	for i := 0; i < 10; i++ {
		tmp[i] = &message{rand.Int()}
	}
	q.AddSlice(tmp)
	if q.Len() != 110 {
		t.Fatalf("AddSlice() fails, wrong number of elements.")
	}

	x, _ := heap.Pop(q).(pq.Queueable)
	for _, v := range q.Collection() {
		if x.Priority() > v.Priority() {
			t.Fatalf("AddSlice() fails, Q not min-heaped.")
		}
	}
}

func TestPopElement(t *testing.T) {
	q := pq.NewQueue()
	q.Push(&message{128})
	q.Push(&message{256})
	q.Push(&message{512})

	x, ok := q.Pop().(pq.Queueable)
	if !ok {
		t.Fatalf("Pop() fails.")
	}
	if x.Priority() != 128 && q.Len() != 2 {
		t.Fatalf("Pop() fails, wrong element or Q is wrong size.")
	}
}

func TestGetElement(t *testing.T) {
	q := initCollection()
	for i := 0; i < 100; i++ {
		if _, err := q.Member(i); err != nil {
			t.Fatalf("Member() fails, wrong number of elements.")
		}
	}
}

func TestOutOfBounds(t *testing.T) {
	q := initCollection()
	_, err := q.Member(100)
	if err != pq.QERR_OUT_OF_BOUNDS {
		t.Fatalf("Member() fails, out of bounds not detected.", q.Len())
	}
}

func TestReturnCollection(t *testing.T) {
	q := initCollection()
	c := q.Collection()
	if len(c) != 100 {
		t.Fatalf("Collection() fails, wrong number of elements.")
	}
	for _, v := range q.Collection() {
		if v, ok := v.(pq.Queueable); !ok {
			t.Fatalf("Collection() fails", v)
		}
	}
}

func TestHeapInitAndPop(t *testing.T) {
	q := initSmallCollection(50)

	heap.Init(q)

	x, _ := heap.Pop(q).(pq.Queueable)

	for _, y := range q.Collection() {
		if x.Priority() > y.Priority() {
			t.Fatalf("heap.Pop() fails, not min element")
		}
	}
}

func TestHeapInitAndPush(t *testing.T) {
	q := initSmallCollection(200)
	heap.Init(q)

	heap.Push(q, &message{0})

	if q.Len() != 11 {
		t.Fatalf("heap.Push() fails, q is wrong size")
	}

	m, err := q.Member(0)
	if err != nil {
		t.Fatalf("Error ¬_¬")
	}
	if m.Priority() != 0 {
		t.Fatalf("heap.Push() fails, pushed valued not found at bottom of heap.")
	}
}

func TestHeapRemove(t *testing.T) {
	q := pq.NewQueue()
	q.Push(&message{8})
	q.Push(&message{3})
	q.Push(&message{9})
	q.Push(&message{1})
	q.Push(&message{10})
	q.Push(&message{2})
	q.Push(&message{5})
	q.Push(&message{6})
	q.Push(&message{7})
	q.Push(&message{4})
	heap.Init(q)

	r := heap.Remove(q, 5).(pq.Queueable)
	fmt.Println("removed:", r)

	if r.Priority() != 9 {
		t.Fatalf("Remove() fails, incorrect member removed.", r.Priority())
	}

	mb, _ := q.Member(0)
	if mb.Priority() != 1 {
		t.Fatalf("Remove() fails, unexpected value for min member")
	}
}
