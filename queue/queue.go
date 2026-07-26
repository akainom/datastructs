package queue

import (
	"datastructs/compare"
	err "datastructs/errors"
	"sync"
)

type Queue[T any] struct {
	m       *sync.RWMutex
	arr     []T
	compare func(a, b T) bool
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		m:       &sync.RWMutex{},
		arr:     make([]T, 0),
		compare: compare.DefaultCompare[T](),
	}
}

func NewQueueFromSlice[T any](s []T) *Queue[T] {
	maxSize := 17359360 // 135620 KB for Queue[int]
	if len(s) > maxSize {
		s = s[:maxSize]
	}

	q := NewQueue[T]()
	q.arr = s

	return q
}

func (q *Queue[T]) Push(val T) error {
	q.m.Lock()
	defer q.m.Unlock()

	if len(q.arr) > 17359359 { // 135620 KB for Queue[int]
		return err.ErrorOverflow
	}

	q.arr = append(q.arr, val)
	return nil
}

func (q *Queue[T]) Pop() (T, error) {
	q.m.Lock()
	defer q.m.Unlock()

	if len(q.arr) == 0 {
		var zero T
		return zero, err.ErrorEmpty
	}

	val := q.arr[0]
	q.arr = q.arr[1:]
	return val, nil
}

func (q *Queue[T]) Front() (T, error) {
	q.m.RLock()
	defer q.m.RUnlock()

	if len(q.arr) == 0 {
		var zero T
		return zero, err.ErrorEmpty
	}

	return q.arr[0], nil
}

func (q *Queue[T]) Len() int {
	q.m.RLock()
	defer q.m.RUnlock()

	return len(q.arr)
}

func (q *Queue[T]) IsEmpty() bool {
	q.m.RLock()
	defer q.m.RUnlock()

	return len(q.arr) == 0
}

func (q *Queue[T]) Exists(target T) bool {
	q.m.RLock()
	defer q.m.RUnlock()

	for _, current := range q.arr {
		if q.compare(current, target) {
			return true
		}
	}

	return false
}

func (q *Queue[T]) Clear() {
	q.m.Lock()
	defer q.m.Unlock()

	q.arr = make([]T, 0)
}

func (q *Queue[T]) ToSlice() []T {
	q.m.RLock()
	defer q.m.RUnlock()

	slice := make([]T, len(q.arr))
	copy(slice, q.arr)

	return slice
}

func (q *Queue[T]) SetCompareFunc(cmp func(a, b T) bool) {
	q.m.Lock()
	defer q.m.Unlock()

	q.compare = cmp
}
