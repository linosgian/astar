package main

import (
	"container/heap"
	"testing"
)

func TestPush(t *testing.T) {
	var pqPushTests = []struct {
		name string
		n    *Node
	}{
		{"test normal case", &Node{}},
	}
	for _, tc := range pqPushTests {
		t.Run(tc.name, func(t *testing.T) {
			pq := PriorityQueue{}
			heap.Init(&pq)

			pq.Push(tc.n)
			if pq[len(pq)-1] != tc.n {
				t.Errorf("Testcase: %q: last element on heap is not %+v", tc.name, *tc.n)
			}
		})
	}
}

func TestPop(t *testing.T) {
	var pqPopTests = []struct {
		name string
		n    *Node
	}{
		{"test normal case", &Node{}},
	}
	for _, tc := range pqPopTests {
		t.Run(tc.name, func(t *testing.T) {
			pq := PriorityQueue{}
			heap.Init(&pq)

			pq.Push(tc.n)
			if pq.Pop() != tc.n {
				t.Errorf("Testcase: %q: expected the inserted node %+v", tc.name, *tc.n)
			}
		})
	}
}
