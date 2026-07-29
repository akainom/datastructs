package list

import (
	"container/list"
	"testing"
)

func BenchmarkAppendList(b *testing.B) {
	l := NewList[any]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Append(i)
	}
}

func BenchmarkAppendContainer(b *testing.B) {
	l := list.New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.PushBack(i)
	}
}

func BenchmarkPopList(b *testing.B) {
	l := NewList[any]()
	for i := 0; i < 100000; i++ {
		l.Append(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Pop()
	}
}

func BenchmarkPopContainer(b *testing.B) {
	l := list.New()
	for i := 0; i < 100000; i++ {
		l.PushBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N && i < 100000; i++ {
		e := l.Back()
		if e != nil {
			l.Remove(e)
		}
	}
}
