package stack

import (
	cmp "datastructs/compare"
	err "datastructs/errors"
	"sync"
)

type Stack[T any] struct {
	arr     []T
	m       *sync.RWMutex
	compare func(a, b T) bool
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		arr:     make([]T, 0),
		m:       &sync.RWMutex{},
		compare: cmp.DefaultCompare[T](),
	}
}

func NewStackFromSlice[T any](s []T) *Stack[T] {
	stack := NewStack[T]()
	maxSize := 4294967295
	if len(s) > maxSize {
		s = s[:maxSize]
	}

	stack.arr = s
	return stack
}

func (s *Stack[T]) Push(val T) error {
	s.m.Lock()
	defer s.m.Unlock()

	if len(s.arr) > 4294967294 {
		return err.ErrorOverflow
	}

	s.arr = append(s.arr, val)
	return nil
}

func (s *Stack[T]) Top() (T, error) {
	s.m.RLock()
	defer s.m.RUnlock()

	if len(s.arr) == 0 {
		var zero T
		return zero, err.ErrorEmpty
	}
	return s.arr[len(s.arr)-1], nil
}

func (s *Stack[T]) Pop() (T, error) {
	s.m.Lock()
	defer s.m.Unlock()

	if len(s.arr) == 0 {
		var zero T
		return zero, err.ErrorEmpty
	}

	popped := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]

	return popped, nil
}

func (s *Stack[T]) IsEmpty() bool {
	s.m.RLock()
	defer s.m.RUnlock()

	return len(s.arr) == 0
}

func (s *Stack[T]) Exists(target T) bool {
	s.m.RLock()
	defer s.m.RUnlock()

	for _, current := range s.arr {
		if s.compare(current, target) {
			return true
		}
	}

	return false
}

func (s *Stack[T]) SetTop(value T) error {
	s.m.Lock()
	defer s.m.Unlock()

	if len(s.arr) == 0 {
		return err.ErrorEmpty
	}

	s.arr[len(s.arr)-1] = value
	return nil
}

func (s *Stack[T]) Len() int {
	s.m.RLock()
	defer s.m.RUnlock()

	return len(s.arr)
}

func (s *Stack[T]) Clear() {
	s.m.Lock()
	defer s.m.Unlock()

	s.arr = make([]T, 0)
}

func (s *Stack[T]) ToSlice() []T {
	s.m.RLock()
	defer s.m.RUnlock()

	slice := make([]T, len(s.arr))
	copy(slice, s.arr)

	return slice
}

func (s *Stack[T]) SetCompareFunc(cmp func(a, b T) bool) {
	s.m.Lock()
	defer s.m.Unlock()

	s.compare = cmp
}
