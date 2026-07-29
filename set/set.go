package set

import (
	err "datastructs/errors"
	"sync"
)

type Set[T comparable] struct {
	m     *sync.RWMutex
	items map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m:     &sync.RWMutex{},
		items: make(map[T]struct{}),
	}
}

func NewSetFromSlice[T comparable](s []T) *Set[T] {
	set := NewSet[T]()
	maxSize := 17359360
	if len(s) > maxSize {
		s = s[:maxSize]
	}
	for _, val := range s {
		set.Add(val)
	}

	return set
}

func (s *Set[T]) Add(val T) error {
	s.m.Lock()
	defer s.m.Unlock()

	if len(s.items) > 17359359 {
		return err.ErrorOverflow
	}

	s.items[val] = struct{}{}
	return nil
}

func (s *Set[T]) Exists(val T) bool {
	s.m.RLock()
	defer s.m.RUnlock()

	_, ok := s.items[val]
	return ok
}

func (s *Set[T]) Remove(val T) error {
	s.m.Lock()
	defer s.m.Unlock()

	if _, ok := s.items[val]; !ok {
		return err.ErrorNF
	}

	delete(s.items, val)
	return nil
}

func (s *Set[T]) Len() int {
	s.m.RLock()
	defer s.m.RUnlock()

	return len(s.items)
}

func (s *Set[T]) Clear() {
	s.m.Lock()
	defer s.m.Unlock()

	s.items = make(map[T]struct{})
}

func (s *Set[T]) IsEmpty() bool {
	s.m.RLock()
	defer s.m.RUnlock()

	return len(s.items) == 0
}

func (s *Set[T]) ToSlice() []T {
	s.m.RLock()
	defer s.m.RUnlock()

	slice := make([]T, 0, len(s.items))
	for val := range s.items {
		slice = append(slice, val)
	}

	return slice
}
