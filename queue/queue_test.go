package queue

import (
	"datastructs/errors"
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	s := NewQueue[any]()

	if s == nil {
		t.Error("NewQueue() returned nil")
	}

	if s.arr == nil {
		t.Error("Expected s.arr to be initialized, got nil")
	}

	if len(s.arr) != 0 {
		t.Errorf("Expected len(s.arr) to be 0, got %d", len(s.arr))
	}

	if s.m == nil {
		t.Error("Expected s.m to be initialized, got nil")
	}
}

func TestNewQueueFromSlice(t *testing.T) {
	slice := []int{4, 3, 2, 1}
	q := NewQueueFromSlice(slice)

	if q == nil {
		t.Error("NewQueueFromSlice() returned nil")
	}

	if len(q.arr) != 4 {
		t.Errorf("Expected len(q.arr) to be 4, got %d", len(q.arr))
	}

	if val, _ := q.Front(); val != 4 {
		t.Errorf("Expected q.Front() to be 4, got %v", val)
	}

	for i, exp := range slice {
		val, err := q.Pop()
		if err != nil {
			t.Errorf("At index %d got error: %v", i, err)
		}

		if !reflect.DeepEqual(val, exp) {
			t.Errorf("At index %d: expected %v, got %v", i, exp, val)
		}
	}

	if len(q.arr) != 0 {
		t.Errorf("Expected len(q.arr) to be 0 after all pops, got %d", len(q.arr))
	}

	strslice := []string{"c", "b", "a"}
	strq := NewQueueFromSlice(strslice)

	for i, exp := range strslice {
		top, err := strq.Pop()
		if err != nil {
			t.Errorf("At index %d got error: %v", i, err)
		}

		if !reflect.DeepEqual(top, exp) {
			t.Errorf("At index %d: expected %v, got %v", i, exp, top)
		}
	}

	empty := NewQueueFromSlice([]int{})
	if len(empty.arr) != 0 {
		t.Errorf("Expected arr length=0 for empty slice, got %d", len(empty.arr))
	}

	bigSlice := make([]int, 100000)
	for i := range bigSlice {
		bigSlice[i] = i
	}
	bigQueue := NewQueueFromSlice(bigSlice)
	if len(bigQueue.arr) != 65535 {
		t.Errorf("Expected arr length=65535 (max), got %d", len(bigQueue.arr))
	}
	if bigQueue.arr[0] != 0 {
		t.Errorf("Expected first element 0, got %v", bigQueue.arr[0])
	}
	if bigQueue.arr[65534] != 65534 {
		t.Errorf("Expected last element 65534, got %v", bigQueue.arr[65534])
	}
}

func TestPush(t *testing.T) {
	s := NewQueue[int]()

	s.Push(42)
	if len(s.arr) != 1 {
		t.Errorf("Expected arr length=1, got %d", len(s.arr))
	}
	if s.arr[0] != 42 {
		t.Errorf("Expected arr[0]=42, got %v", s.arr[0])
	}

	s.Push(100)
	if len(s.arr) != 2 {
		t.Errorf("Expected arr length=2, got %d", len(s.arr))
	}
	if s.arr[1] != 100 {
		t.Errorf("Expected arr[1]=100, got %v", s.arr[1])
	}

	s.Push(200)
	if len(s.arr) != 3 {
		t.Errorf("Expected arr length=3, got %d", len(s.arr))
	}
	if s.arr[2] != 200 {
		t.Errorf("Expected arr[2]=200, got %v", s.arr[2])
	}
}

func TestPushOverflow(t *testing.T) {
	q := NewQueue[int]()

	for i := 0; i < 65535; i++ {
		err := q.Push(i)
		if err != nil {
			t.Errorf("Push(%d) unexpected error: %v", i, err)
		}
	}

	if len(q.arr) != 65535 {
		t.Errorf("Expected len(q.arr)=65535, got %d", len(q.arr))
	}

	err := q.Push(65535)
	if err != errors.ErrorOverflow {
		t.Errorf("Expected ErrorOverflow, got %v", err)
	}
	if len(q.arr) != 65535 {
		t.Errorf("Expected len(q.arr) still 65535, got %d", len(q.arr))
	}
}

func TestPop(t *testing.T) {
	s := NewQueueFromSlice([]int{1, 2, 3, 4})

	expected := []int{1, 2, 3, 4}
	for i, exp := range expected {
		val, err := s.Pop()
		if err != nil {
			t.Errorf("At index %d got unexpected error: %v", i, err)
		}
		if val != exp {
			t.Errorf("At index %d: expected %d, got %d", i, exp, val)
		}
	}

	if len(s.arr) != 0 {
		t.Errorf("Expected arr length=0 after all pops, got %d", len(s.arr))
	}

	val, err := s.Pop()
	if err != errors.ErrorEmpty {
		t.Errorf("Expected ErrorEmpty, got %v", err)
	}
	var zero int
	if val != zero {
		t.Errorf("Expected zero value (0), got %v", val)
	}
}

func TestFront(t *testing.T) {
	q := NewQueueFromSlice([]int{1, 2, 3})

	val, err := q.Front()
	if err != nil {
		t.Errorf("Front() unexpected error: %v", err)
	}

	if val != 1 {
		t.Errorf("Expected Front()=1, got %v", val)
	}

	if len(q.arr) != 3 {
		t.Errorf("Expected arr length=3 after Front(), got %d", len(q.arr))
	}

	popped, _ := q.Pop()
	if popped != 1 {
		t.Errorf("Expected Pop()=1 after Front(), got %v", popped)
	}

	empty := NewQueue[int]()
	val, err = empty.Front()
	if err != errors.ErrorEmpty {
		t.Errorf("Expected ErrorEmpty from empty queue, got %v", err)
	}
}

func TestExists(t *testing.T) {
	q := NewQueueFromSlice([]any{1, "hello", 3.14, []int{1, 2}})

	if !q.Exists("hello") {
		t.Error("Expected Exists('hello') to be true")
	}

	if !q.Exists(1) {
		t.Error("Expected Exists(1) to be true")
	}

	if !q.Exists(3.14) {
		t.Error("Expected Exists(3.14) to be true")
	}

	if !q.Exists([]int{1, 2}) {
		t.Error("Expected Exists([1 2]) to be true")
	}

	if q.Exists("nonexistent") {
		t.Error("Expected Exists('nonexistent') to be false")
	}

	if q.Exists(999) {
		t.Error("Expected Exists(999) to be false")
	}

	if q.Exists([]int{3, 4}) {
		t.Error("Expected Exists([3 4]) to be false")
	}

	empty := NewQueue[int]()
	if empty.Exists(42) {
		t.Error("Expected Exists(42) to be false for empty queue")
	}
}

func TestLen(t *testing.T) {
	q := NewQueue[int]()

	if q.Len() != 0 {
		t.Errorf("Expected len=0, got %d", q.Len())
	}

	q.Push(1)
	if q.Len() != 1 {
		t.Errorf("Expected len=1, got %d", q.Len())
	}

	q.Push(2)
	q.Push(3)
	if q.Len() != 3 {
		t.Errorf("Expected len=3, got %d", q.Len())
	}

	q.Pop()
	if q.Len() != 2 {
		t.Errorf("Expected len=2 after Pop, got %d", q.Len())
	}

	q.Clear()
	if q.Len() != 0 {
		t.Errorf("Expected len=0 after Clear, got %d", q.Len())
	}
}

func TestIsEmpty(t *testing.T) {
	q := NewQueue[int]()

	if !q.IsEmpty() {
		t.Error("Expected IsEmpty() to be true for new queue")
	}

	q.Push(1)
	if q.IsEmpty() {
		t.Error("Expected IsEmpty() to be false after Push()")
	}

	q.Pop()
	if !q.IsEmpty() {
		t.Error("Expected IsEmpty() to be true after Pop()")
	}

	q.Push(1)
	q.Push(2)
	q.Clear()
	if !q.IsEmpty() {
		t.Error("Expected IsEmpty() to be true after Clear")
	}
}

func TestClear(t *testing.T) {
	q := NewQueueFromSlice([]int{1, 2, 3, 4, 5})

	if len(q.arr) != 5 {
		t.Errorf("Expected arr length=5 before Clear, got %d", len(q.arr))
	}

	q.Clear()

	if len(q.arr) != 0 {
		t.Errorf("Expected arr length=0 after Clear, got %d", len(q.arr))
	}
	if q.arr == nil {
		t.Error("Expected arr to be initialized (empty slice), got nil")
	}

	_, err := q.Pop()
	if err != errors.ErrorEmpty {
		t.Errorf("Expected ErrorEmpty after Clear, got %v", err)
	}

	_, err = q.Front()
	if err != errors.ErrorEmpty {
		t.Errorf("Expected ErrorEmpty after Clear, got %v", err)
	}
}

func TestToSlice(t *testing.T) {
	empty := NewQueue[int]()
	result := empty.ToSlice()
	if len(result) != 0 {
		t.Errorf("Expected empty slice, got %v", result)
	}

	q := NewQueueFromSlice([]int{1, 2, 3, 4})
	result = q.ToSlice()

	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	if len(q.arr) != 4 {
		t.Errorf("Expected arr length=4 after ToSlice, got %d", len(q.arr))
	}

	q.Push(5)
	if len(result) != 4 {
		t.Errorf("Expected result length still 4, got %d", len(result))
	}
}

func TestConcurrentOperations(t *testing.T) {
	q := NewQueue[int]()
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(val int) {
			q.Push(val)
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	if q.Len() != 10 {
		t.Errorf("Expected len=10, got %d", q.Len())
	}

	for i := 0; i < 10; i++ {
		go func() {
			q.Pop()
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	if q.Len() != 0 {
		t.Errorf("Expected len=0 after all pops, got %d", q.Len())
	}
}

func BenchmarkPush(b *testing.B) {
	q := NewQueue[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	q := NewQueue[int]()
	for i := 0; i < 100000; i++ {
		q.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func BenchmarkFront(b *testing.B) {
	q := NewQueue[int]()
	for i := 0; i < 100000; i++ {
		q.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Front()
	}
}
