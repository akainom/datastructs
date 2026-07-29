package pqueue

import (
	"container/heap"
	"testing"
)

func BenchmarkPQueuePush(b *testing.B) {
	pq := NewPQueue[any]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Push(i, i)
	}
}

func BenchmarkPQueuePop(b *testing.B) {
	pq := NewPQueue[any]()
	for i := 0; i < 100000; i++ {
		pq.Push(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Pop()
	}
}

type HeapItem struct {
	Value    int
	Priority int
	Index    int
}

type Heap []HeapItem

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].Priority > h[j].Priority }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i]; h[i].Index = i; h[j].Index = j }

func (h *Heap) Push(x any) {
	n := len(*h)
	item := x.(HeapItem)
	item.Index = n
	*h = append(*h, item)
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = HeapItem{}
	*h = old[0 : n-1]
	return item
}

func BenchmarkHeapPush(b *testing.B) {
	h := &Heap{}
	heap.Init(h)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Push(h, HeapItem{Value: i, Priority: i})
	}
}

func BenchmarkHeapPop(b *testing.B) {
	h := &Heap{}
	heap.Init(h)

	for i := 0; i < b.N; i++ {
		heap.Push(h, HeapItem{Value: i, Priority: i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Pop(h)
	}
}
