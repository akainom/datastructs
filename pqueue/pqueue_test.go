package pqueue

import (
	"datastructs/errors"

	"reflect"
	"testing"
)

func TestNewPQueue(t *testing.T) {
	pq := NewPQueue[int]()

	if pq == nil {
		t.Error("NewPQueue() returned nil")
	}

	if pq.arr == nil {
		t.Error("Expected pq.arr to be initialized, got nil")
	}

	if len(pq.arr) != 0 {
		t.Errorf("Expected len(pq.arr) to be 0, got %d", len(pq.arr))
	}

	if pq.m == nil {
		t.Error("Expected pq.m to be initialized, got nil")
	}
}

func TestPush(t *testing.T) {
	pq := NewPQueue[int]()

	err := pq.Push(10, 5)
	if err != nil {
		t.Errorf("Push() unexpected error: %v", err)
	}
	if len(pq.arr) != 1 {
		t.Errorf("Expected len=1, got %d", len(pq.arr))
	}
	if pq.arr[0].Val != 10 {
		t.Errorf("Expected arr[0].Val=10, got %v", pq.arr[0].Val)
	}
	if pq.arr[0].Priority != 5 {
		t.Errorf("Expected arr[0].Priority=5, got %d", pq.arr[0].Priority)
	}

	pq.Push(20, 10)
	if len(pq.arr) != 2 {
		t.Errorf("Expected len=2, got %d", len(pq.arr))
	}
	if pq.arr[0].Val != 20 {
		t.Errorf("Expected корень 20, got %v", pq.arr[0].Val)
	}
	if pq.arr[0].Priority != 10 {
		t.Errorf("Expected корень приоритет 10, got %d", pq.arr[0].Priority)
	}

	pq.Push(30, 7)
	if len(pq.arr) != 3 {
		t.Errorf("Expected len=3, got %d", len(pq.arr))
	}
	if pq.arr[0].Val != 20 {
		t.Errorf("Expected корень 20, got %v", pq.arr[0].Val)
	}
}

func TestPushOverflow(t *testing.T) {
	pq := NewPQueue[int]()

	for i := 0; i < 8679680; i++ {
		err := pq.Push(i, i)
		if err != nil {
			t.Errorf("Push(%d) unexpected error: %v", i, err)
		}
	}

	if len(pq.arr) != 8679680 {
		t.Errorf("Expected len=8679680, got %d", len(pq.arr))
	}

	err := pq.Push(999, 999)
	if err != errors.ErrorOverflow {
		t.Errorf("Expected ErrorOverflow, got %v", err)
	}
}

func TestPop(t *testing.T) {
	pq := NewPQueue[int]()
	pq.Push(10, 1)
	pq.Push(20, 5)
	pq.Push(30, 3)
	pq.Push(40, 10)
	pq.Push(50, 7)

	expected := []int{40, 50, 20, 30, 10}
	for _, exp := range expected {
		val, err := pq.Pop()
		if err != nil {
			t.Errorf("Pop() unexpected error: %v", err)
		}
		if val != exp {
			t.Errorf("Expected %d, got %d", exp, val)
		}
	}

	if len(pq.arr) != 0 {
		t.Errorf("Expected len=0 after all pops, got %d", len(pq.arr))
	}

	_, err := pq.Pop()
	if err != errors.ErrorEmpty {
		t.Errorf("Expected ErrorEmpty, got %v", err)
	}
}

func TestPopWithEqualPriority(t *testing.T) {
	pq := NewPQueue[string]()

	pq.Push("first", 1)
	pq.Push("second", 1)
	pq.Push("third", 1)

	vals := make([]string, 0, 3)
	for i := 0; i < 3; i++ {
		val, err := pq.Pop()
		if err != nil {
			t.Errorf("Pop() unexpected error: %v", err)
		}
		vals = append(vals, val)
	}

	expected := map[string]bool{"first": true, "second": true, "third": true}
	for _, v := range vals {
		if !expected[v] {
			t.Errorf("Unexpected value: %s", v)
		}
		delete(expected, v)
	}
}

func TestPeek(t *testing.T) {
	pq := NewPQueue[int]()
	pq.Push(10, 1)
	pq.Push(20, 5)
	pq.Push(30, 3)

	val, err := pq.Peek()
	if err != nil {
		t.Errorf("Peek() unexpected error: %v", err)
	}
	if val != 20 {
		t.Errorf("Expected Peek()=20, got %v", val)
	}

	if len(pq.arr) != 3 {
		t.Errorf("Expected len=3 after Peek, got %d", len(pq.arr))
	}

	empty := NewPQueue[int]()
	_, err = empty.Peek()
	if err != errors.ErrorEmpty {
		t.Errorf("Expected ErrorEmpty from empty queue, got %v", err)
	}
}

func TestLen(t *testing.T) {
	pq := NewPQueue[int]()

	if pq.Len() != 0 {
		t.Errorf("Expected len=0, got %d", pq.Len())
	}

	pq.Push(1, 1)
	if pq.Len() != 1 {
		t.Errorf("Expected len=1, got %d", pq.Len())
	}

	pq.Push(2, 2)
	pq.Push(3, 3)
	if pq.Len() != 3 {
		t.Errorf("Expected len=3, got %d", pq.Len())
	}

	pq.Pop()
	if pq.Len() != 2 {
		t.Errorf("Expected len=2 after Pop, got %d", pq.Len())
	}

	pq.Clear()
	if pq.Len() != 0 {
		t.Errorf("Expected len=0 after Clear, got %d", pq.Len())
	}
}

func TestIsEmpty(t *testing.T) {
	pq := NewPQueue[int]()

	if !pq.IsEmpty() {
		t.Error("Expected IsEmpty() to be true for new queue")
	}

	pq.Push(1, 1)
	if pq.IsEmpty() {
		t.Error("Expected IsEmpty() to be false after Push")
	}

	pq.Pop()
	if !pq.IsEmpty() {
		t.Error("Expected IsEmpty() to be true after Pop")
	}

	pq.Push(1, 1)
	pq.Push(2, 2)
	pq.Clear()
	if !pq.IsEmpty() {
		t.Error("Expected IsEmpty() to be true after Clear")
	}
}

func TestClear(t *testing.T) {
	pq := NewPQueue[int]()
	pq.Push(1, 1)
	pq.Push(2, 2)
	pq.Push(3, 3)
	pq.Push(4, 4)
	pq.Push(5, 5)

	if len(pq.arr) != 5 {
		t.Errorf("Expected len=5 before Clear, got %d", len(pq.arr))
	}

	pq.Clear()

	if len(pq.arr) != 0 {
		t.Errorf("Expected len=0 after Clear, got %d", len(pq.arr))
	}
	if pq.arr == nil {
		t.Error("Expected arr to be initialized (empty slice), got nil")
	}

	_, err := pq.Pop()
	if err != errors.ErrorEmpty {
		t.Errorf("Expected ErrorEmpty after Clear, got %v", err)
	}
}

func TestToSlice(t *testing.T) {
	pq := NewPQueue[int]()

	empty := pq.ToSlice()
	if len(empty) != 0 {
		t.Errorf("Expected empty slice, got %v", empty)
	}

	pq.Push(10, 1)
	pq.Push(20, 5)
	pq.Push(30, 3)

	result := pq.ToSlice()
	expected := []PNode[int]{
		{Val: 20, Priority: 5},
		{Val: 10, Priority: 1},
		{Val: 30, Priority: 3},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestToSliceOfValues(t *testing.T) {
	pq := NewPQueue[int]()

	empty := pq.ToSliceOfValues()
	if len(empty) != 0 {
		t.Errorf("Expected empty slice, got %v", empty)
	}

	pq.Push(10, 1)
	pq.Push(20, 5)
	pq.Push(30, 3)

	result := pq.ToSliceOfValues()
	expected := []int{20, 10, 30}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestConcurrentOperations(t *testing.T) {
	pq := NewPQueue[int]()
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(val int) {
			pq.Push(val, val)
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	if pq.Len() != 10 {
		t.Errorf("Expected len=10, got %d", pq.Len())
	}

	for i := 0; i < 10; i++ {
		go func() {
			pq.Pop()
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	if pq.Len() != 0 {
		t.Errorf("Expected len=0 after all pops, got %d", pq.Len())
	}
}

func TestConcurrentPushOverflow(t *testing.T) {
	pq := NewPQueue[int]()
	done := make(chan bool)

	for i := 0; i < 8679680; i++ {
		err := pq.Push(i, i)
		if err != nil {
			t.Errorf("Push(%d) unexpected error: %v", i, err)
		}
	}

	for i := 0; i < 10; i++ {
		go func() {
			err := pq.Push(999, 999)
			if err != errors.ErrorOverflow {
				t.Errorf("Expected ErrorOverflow, got %v", err)
			}
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	if pq.Len() != 8679680 {
		t.Errorf("Expected len still %d, got %d", 8679679, pq.Len())
	}
}

func BenchmarkPush(b *testing.B) {
	pq := NewPQueue[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Push(i, i)
	}
}

func BenchmarkPop(b *testing.B) {
	pq := NewPQueue[int]()
	for i := 0; i < 100000; i++ {
		pq.Push(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Pop()
	}
}

func BenchmarkPeek(b *testing.B) {
	pq := NewPQueue[int]()
	for i := 0; i < 100000; i++ {
		pq.Push(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Peek()
	}
}
