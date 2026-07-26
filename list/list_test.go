package list

import (
	"datastructs/errors"
	"reflect"
	"testing"
)

func TestNewList(t *testing.T) {
	l := NewList[any]()

	if l == nil {
		t.Error("NewList() returned nil")
	}

	if l.Head != nil {
		t.Errorf("Expected Head to be nil, got %v", l.Head)
	}

	if l.Tail != nil {
		t.Errorf("Expected Tail to be nil, got %v", l.Tail)
	}

	if l.len != 0 {
		t.Errorf("Expected len to be 0, got %d", l.len)
	}
}

func TestAppend(t *testing.T) {
	l := NewList[any]()

	l.Append(42)
	if l.len != 1 {
		t.Errorf("Expected len=1, got %d", l.len)
	}
	if l.Head == nil || l.Head.Val != 42 {
		t.Errorf("Expected Head.Val=42, got %v", l.Head)
	}
	if l.Tail == nil || l.Tail.Val != 42 {
		t.Errorf("Expected Tail.Val=42, got %v", l.Tail)
	}

	l.Append("hello")
	if l.len != 2 {
		t.Errorf("Expected len=2, got %d", l.len)
	}
	if l.Tail.Val != "hello" {
		t.Errorf("Expected Tail.Val='hello', got %v", l.Tail.Val)
	}
	if l.Head.Next.Val != "hello" {
		t.Errorf("Expected Head.Next.Val='hello', got %v", l.Head.Next.Val)
	}
	if l.Tail.Prev.Val != 42 {
		t.Errorf("Expected Tail.Prev.Val=42, got %v", l.Tail.Prev.Val)
	}

	l.Append(3.14)
	if l.len != 3 {
		t.Errorf("Expected len=3, got %d", l.len)
	}
	if l.Tail.Val != 3.14 {
		t.Errorf("Expected Tail.Val=3.14, got %v", l.Tail.Val)
	}

	l.Clear()
	for range 65535 {
		l.Append(1)
	}

	err := l.Append(1)
	if err != errors.ErrorOverflow {
		t.Errorf("Expected overflow error, got %v", err)
	}
}

func TestNewListFromSlice(t *testing.T) {
	l1 := NewListFromSlice([]any{})
	if l1.len != 0 {
		t.Errorf("Expected len=0, got %d", l1.len)
	}

	input := []any{1, "two", 3.0}
	l2 := NewListFromSlice(input)

	if l2.len != 3 {
		t.Errorf("Expected len=3, got %d", l2.len)
	}

	expected := []any{1, "two", 3.0}
	current := l2.Head
	for i, exp := range expected {
		if current == nil {
			t.Errorf("List is shorter than expected at index %d", i)
			break
		}
		if !reflect.DeepEqual(current.Val, exp) {
			t.Errorf("At index %d: expected %v, got %v", i, exp, current.Val)
		}
		current = current.Next
	}

	overflow := make([]int, 0, 65536)
	for i := range 65536 {
		overflow = append(overflow, i)
	}

	l3 := NewListFromSlice(overflow)
	if l3.len != 65535 {
		t.Errorf("Expected len to be max(65535), got %v", l3.len)
	}
}

func TestPop(t *testing.T) {
	l := NewListFromSlice([]any{1, 2, 3})

	node, err := l.Pop()
	if err != nil {
		t.Errorf("Pop() unexpected error: %v", err)
	}
	if node.Val != 3 {
		t.Errorf("Expected popped value 3, got %v", node.Val)
	}
	if l.len != 2 {
		t.Errorf("Expected len=2, got %d", l.len)
	}
	if l.Tail.Val != 2 {
		t.Errorf("Expected new Tail=2, got %v", l.Tail.Val)
	}

	node, _ = l.Pop()
	if node.Val != 2 {
		t.Errorf("Expected popped value 2, got %v", node.Val)
	}
	if l.len != 1 {
		t.Errorf("Expected len=1, got %d", l.len)
	}
	if l.Head.Val != 1 || l.Tail.Val != 1 {
		t.Errorf("Expected Head=Tail=1, got Head=%v, Tail=%v", l.Head.Val, l.Tail.Val)
	}

	node, _ = l.Pop()
	if node.Val != 1 {
		t.Errorf("Expected popped value 1, got %v", node.Val)
	}
	if l.len != 0 {
		t.Errorf("Expected len=0, got %d", l.len)
	}
	if l.Head != nil || l.Tail != nil {
		t.Errorf("Expected Head=Tail=nil, got Head=%v, Tail=%v", l.Head, l.Tail)
	}

	_, err = l.Pop()
	if err != errors.ErrorEmpty {
		t.Errorf("Expected ErrorEmpty, got %v", err)
	}
}

func TestGet(t *testing.T) {
	l := NewListFromSlice([]any{42, "hello", 3.14, []int{1, 2}})

	node, err := l.Get("hello")
	if err != nil {
		t.Errorf("Get() unexpected error: %v", err)
	}
	if node.Val != "hello" {
		t.Errorf("Expected 'hello', got %v", node.Val)
	}

	node, err = l.Get(42)
	if err != nil {
		t.Errorf("Get() unexpected error: %v", err)
	}
	if node.Val != 42 {
		t.Errorf("Expected 42, got %v", node.Val)
	}

	node, err = l.Get([]int{1, 2})
	if err != nil {
		t.Errorf("Get() unexpected error: %v", err)
	}
	if !reflect.DeepEqual(node.Val, []int{1, 2}) {
		t.Errorf("Expected [1 2], got %v", node.Val)
	}

	_, err = l.Get("nonexistent")
	if err != errors.ErrorNF {
		t.Errorf("Expected ErrorNF, got %v", err)
	}

	empty := NewList[any]()
	_, err = empty.Get(42)
	if err != errors.ErrorNF {
		t.Errorf("Expected ErrorNF from empty list, got %v", err)
	}
}

func TestSet(t *testing.T) {
	l := NewListFromSlice([]any{1, 2, 3, 2, 4})

	count, err := l.Set(2, "two")
	if err != nil {
		t.Errorf("Set() unexpected error: %v", err)
	}
	if count == 0 {
		t.Error("Set() returned 0, expected at least 1")
	}

	count, err = l.Set(999, "nine")
	if err != errors.ErrorNF {
		t.Errorf("Expected ErrorNF, got %v", err)
	}
	if count != 0 {
		t.Errorf("Set() returned %d, expected 0", count)
	}

	empty := NewList[any]()
	_, err = empty.Set(42, 100)
	if err != errors.ErrorEmpty {
		t.Errorf("Expected ErrorEmpty, got %v", err)
	}
}

func TestLast(t *testing.T) {
	l := NewListFromSlice([]any{1, 2, 3})

	node, err := l.Last()
	if err != nil {
		t.Errorf("Last() unexpected error: %v", err)
	}
	if node.Val != 3 {
		t.Errorf("Expected 3, got %v", node.Val)
	}

	empty := NewList[any]()
	_, err = empty.Last()
	if err != errors.ErrorEmpty {
		t.Errorf("Expected ErrorEmpty, got %v", err)
	}
}

func TestLen(t *testing.T) {
	l := NewList[any]()

	if l.Len() != 0 {
		t.Errorf("Expected len=0, got %d", l.Len())
	}

	l.Append(1)
	if l.Len() != 1 {
		t.Errorf("Expected len=1, got %d", l.Len())
	}

	l.Append(2)
	l.Append(3)
	if l.Len() != 3 {
		t.Errorf("Expected len=3, got %d", l.Len())
	}

	l.Pop()
	if l.Len() != 2 {
		t.Errorf("Expected len=2, got %d", l.Len())
	}
}

func TestIsEmpty(t *testing.T) {
	l := NewList[any]()

	if !l.IsEmpty() {
		t.Error("Expected IsEmpty() to be true for new list")
	}

	l.Append(1)
	if l.IsEmpty() {
		t.Error("Expected IsEmpty() to be false after adding element")
	}

	l.Pop()
	if !l.IsEmpty() {
		t.Error("Expected IsEmpty() to be true after removing all elements")
	}
}

func TestClear(t *testing.T) {
	l := NewListFromSlice([]any{1, 2, 3})

	l.Clear()

	if l.len != 0 {
		t.Errorf("Expected len=0, got %d", l.len)
	}
	if l.Head != nil {
		t.Errorf("Expected Head=nil, got %v", l.Head)
	}
	if l.Tail != nil {
		t.Errorf("Expected Tail=nil, got %v", l.Tail)
	}

	_, err := l.Pop()
	if err != errors.ErrorEmpty {
		t.Errorf("Expected ErrorEmpty after Clear(), got %v", err)
	}
}

func TestToSlice(t *testing.T) {
	empty := NewList[any]()
	result := empty.ToSlice()
	if len(result) != 0 {
		t.Errorf("Expected empty slice, got %v", result)
	}

	input := []any{1, "hello", 3.14, true}
	l := NewListFromSlice(input)
	result = l.ToSlice()

	if !reflect.DeepEqual(result, input) {
		t.Errorf("Expected %v, got %v", input, result)
	}

	if l.Len() != uint16(len(input)) {
		t.Errorf("Expected len=%d, got %d", len(input), l.Len())
	}
}

func TestConcurrentOperations(t *testing.T) {
	l := NewList[any]()

	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(val int) {
			l.Append(val)
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	if l.Len() != 10 {
		t.Errorf("Expected len=10, got %d", l.Len())
	}
}

func BenchmarkAppend(b *testing.B) {
	l := NewList[int]()
	for i := 0; i < b.N; i++ {
		l.Append(i)
	}
}

func BenchmarkPop(b *testing.B) {
	l := NewList[int]()
	for i := 0; i < b.N; i++ {
		l.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Pop()
	}
}

func BenchmarkGet(b *testing.B) {
	l := NewList[any]()
	for i := 0; i < 1000; i++ {
		l.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Get(999)
	}
}
