package pqueue

import (
	err "datastructs/errors"
	"sync"
)

type PNode[T any] struct {
	Val      T
	Priority int
}

type PQueue[T any] struct {
	m   *sync.RWMutex
	arr []PNode[T]
}

func NewPQueue[T any]() *PQueue[T] {
	return &PQueue[T]{
		m:   &sync.RWMutex{},
		arr: make([]PNode[T], 0),
	}
}

func (pq *PQueue[T]) Push(value T, priority int) error {
	pq.m.Lock()
	defer pq.m.Unlock()

	if len(pq.arr) > 8679679 { // 135620 KB for PQueue[int]
		return err.ErrorOverflow
	}

	node := PNode[T]{Val: value, Priority: priority}
	pq.arr = append(pq.arr, node)
	pq.up(len(pq.arr) - 1)

	return nil
}

func (pq *PQueue[T]) Pop() (T, error) {
	pq.m.Lock()
	defer pq.m.Unlock()

	l := len(pq.arr)
	if l == 0 {
		var zero T
		return zero, err.ErrorEmpty
	}

	root := pq.arr[0]
	last := pq.arr[l-1]

	pq.arr[0] = last
	pq.arr = pq.arr[:l-1]

	if len(pq.arr) > 0 {
		pq.down(0)
	}

	return root.Val, nil
}

func (pq *PQueue[T]) Peek() (T, error) {
	pq.m.RLock()
	defer pq.m.RUnlock()

	if len(pq.arr) == 0 {
		var zero T
		return zero, err.ErrorEmpty
	}

	return pq.arr[0].Val, nil
}

func (pq *PQueue[T]) Len() int {
	pq.m.RLock()
	defer pq.m.RUnlock()

	return len(pq.arr)
}

func (pq *PQueue[T]) IsEmpty() bool {
	pq.m.RLock()
	defer pq.m.RUnlock()

	return len(pq.arr) == 0
}

func (pq *PQueue[T]) Clear() {
	pq.m.Lock()
	defer pq.m.Unlock()

	pq.arr = make([]PNode[T], 0)
}

func (pq *PQueue[T]) ToSliceOfValues() []T {
	pq.m.RLock()
	defer pq.m.RUnlock()

	if len(pq.arr) == 0 {
		return make([]T, 0)
	}

	slice := make([]T, 0, len(pq.arr))

	for _, node := range pq.arr {
		slice = append(slice, node.Val)
	}

	return slice
}

func (pq *PQueue[T]) ToSlice() []PNode[T] {
	pq.m.RLock()
	defer pq.m.RUnlock()

	slice := make([]PNode[T], len(pq.arr))
	copy(slice, pq.arr)
	return slice
}

func (pq *PQueue[T]) up(idx int) {
	for idx > 0 {
		parentIdx := (idx - 1) / 2

		if pq.arr[parentIdx].Priority < pq.arr[idx].Priority {
			pq.arr[parentIdx], pq.arr[idx] = pq.arr[idx], pq.arr[parentIdx]
		} else {
			break
		}

		idx = parentIdx
	}
}

func (pq *PQueue[T]) down(idx int) {
	l := len(pq.arr)
	for {
		leftIdx := idx*2 + 1
		if leftIdx >= l {
			break
		}

		rightIdx := idx*2 + 2
		largest := leftIdx

		if rightIdx < l && pq.arr[rightIdx].Priority > pq.arr[leftIdx].Priority {
			largest = rightIdx
		}

		if pq.arr[largest].Priority > pq.arr[idx].Priority {
			pq.arr[idx], pq.arr[largest] = pq.arr[largest], pq.arr[idx]
			idx = largest
		} else {
			break
		}
	}
}
