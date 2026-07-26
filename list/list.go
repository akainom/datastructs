package list

import (
	cmp "datastructs/compare"
	err "datastructs/errors"
	"sync"
)

type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
	Prev *ListNode[T]
}

type List[T any] struct {
	m       *sync.RWMutex
	len     uint32
	Head    *ListNode[T]
	Tail    *ListNode[T]
	compare func(a, b T) bool
}

func NewList[T any]() *List[T] {
	return &List[T]{
		m:       &sync.RWMutex{},
		len:     0,
		Head:    nil,
		Tail:    nil,
		compare: cmp.DefaultCompare[T](),
	}
}

func NewListFromSlice[T any](s []T) *List[T] {
	list := NewList[T]()

	maxSize := 8679680 // 135620 KB for List[int]
	if len(s) > maxSize {
		s = s[:maxSize]
	}

	if len(s) == 0 {
		return list
	}

	list.m.Lock()
	defer list.m.Unlock()

	for _, val := range s {
		node := &ListNode[T]{
			Val:  val,
			Prev: nil,
			Next: nil,
		}

		if list.Head == nil {
			list.Head = node
			list.Tail = node
		} else {
			node.Prev = list.Tail
			list.Tail.Next = node
			list.Tail = node
		}
		list.len++
	}

	return list
}

func (l *List[T]) Append(val T) error {
	l.m.Lock()
	defer l.m.Unlock()

	if l.len > 8679680 { // 135620 KB for List[int]
		return err.ErrorOverflow
	}

	node := &ListNode[T]{
		Val:  val,
		Prev: nil,
		Next: nil,
	}

	prevTail := l.Tail

	if l.len < 1 {
		l.Head = node
	} else {
		node.Prev = prevTail
		prevTail.Next = node
	}

	l.Tail = node
	l.len++
	return nil
}

func (l *List[T]) Pop() (*ListNode[T], error) {
	l.m.Lock()
	defer l.m.Unlock()

	if l.len == 0 {
		return nil, err.ErrorEmpty
	}

	tail := l.Tail

	if l.len == 1 {
		l.Head = nil
		l.Tail = nil
	} else {
		l.Tail = tail.Prev
		l.Tail.Next = nil
	}

	l.len--
	return tail, nil
}

func (l *List[T]) Get(target T) (*ListNode[T], error) {
	l.m.RLock()
	defer l.m.RUnlock()

	var current = l.Head
	for current != nil {
		if l.compare(current.Val, target) {
			return current, nil
		}
		current = current.Next
	}

	return nil, err.ErrorNF
}

func (l *List[T]) Set(target T, val T) (int, error) {
	l.m.Lock()
	defer l.m.Unlock()

	if l.len < 1 {
		return 0, err.ErrorEmpty
	}

	var (
		current  = l.Head
		chgCount = 0
	)
	for current != nil {
		if l.compare(current.Val, target) {
			chgCount++
			current.Val = val
		}
		current = current.Next
	}

	if chgCount == 0 {
		return chgCount, err.ErrorNF
	}

	return chgCount, nil
}

func (l *List[T]) Last() (*ListNode[T], error) {
	l.m.RLock()
	defer l.m.RUnlock()

	if l.len < 1 {
		return nil, err.ErrorEmpty
	}

	return l.Tail, nil

}

func (l *List[T]) Len() uint32 {
	l.m.RLock()
	defer l.m.RUnlock()

	return l.len
}

func (l *List[T]) IsEmpty() bool {
	l.m.RLock()
	defer l.m.RUnlock()

	if l.len > 0 {
		return false
	}

	return true
}

func (l *List[T]) Clear() {
	l.m.Lock()
	defer l.m.Unlock()

	l.Head = nil
	l.Tail = nil
	l.len = 0
}

func (l *List[T]) ToSlice() []T {
	l.m.RLock()
	defer l.m.RUnlock()

	slice := make([]T, 0, l.len)
	current := l.Head
	for current != nil {
		slice = append(slice, current.Val)
		current = current.Next
	}

	return slice
}

func (l *List[T]) SetCompareFunc(cmp func(a, b T) bool) {
	l.m.Lock()
	defer l.m.Unlock()

	l.compare = cmp
}
