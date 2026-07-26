package stack

import (
	cmp "datastructs/compare"
	err "datastructs/errors"
	"sync"
)

type Stack[T any] struct {
	len     uint16
	arr     []T
	m       *sync.RWMutex
	compare func(a, b T) bool
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		len:     0,
		arr:     make([]T, 0),
		m:       &sync.RWMutex{},
		compare: cmp.DefaultCompare[T](),
	}
}

func NewStackFromSlice[T any](s []T) *Stack[T] {
	stack := NewStack[T]()
	maxSize := 65535
	if len(s) > maxSize {
		s = s[:maxSize]
	}

	stack.arr = s
	stack.len = uint16(len(s))
	return stack
}

func (s *Stack[T]) Push(val T) error {
	s.m.Lock()
	defer s.m.Unlock()

	if s.len > 65534 {
		return err.ErrorOverflow
	}

	s.arr = append(s.arr, val)
	s.len++
	return nil
}

func (s *Stack[T]) Top() (T, error) {
	s.m.RLock()
	defer s.m.RUnlock()

	if s.len < 1 {
		var zero T
		return zero, err.ErrorEmpty
	}
	return s.arr[s.len-1], nil
}

func (s *Stack[T]) Pop() (T, error) {
	s.m.Lock()
	defer s.m.Unlock()

	if s.len < 1 {
		var zero T
		return zero, err.ErrorEmpty
	}

	popped := s.arr[s.len-1]
	s.arr = s.arr[:s.len-1]

	s.len--
	return popped, nil
}

func (s *Stack[T]) IsEmpty() bool {
	s.m.RLock()
	defer s.m.RUnlock()

	if s.len < 1 {
		return true
	}

	return false
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

	if s.len < 1 {
		return err.ErrorEmpty
	}

	s.arr[s.len-1] = value
	return nil
}

func (s *Stack[T]) Len() uint16 {
	s.m.RLock()
	defer s.m.RUnlock()

	return s.len
}

func (s *Stack[T]) Clear() {
	s.m.Lock()
	defer s.m.Unlock()

	s.arr = make([]T, 0)
	s.len = 0
}

func (s *Stack[T]) ToSlice() []T {
	s.m.RLock()
	defer s.m.RUnlock()

	slice := make([]T, s.len)
	copy(slice, s.arr)

	return slice
}

func (s *Stack[T]) SetCompareFunc(cmp func(a, b T) bool) {
	s.m.Lock()
	defer s.m.Unlock()

	s.compare = cmp
}
