package stack

import (
	"datastructs/errors"
	"math/rand"
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack[any]()

	if s == nil {
		t.Error("NewStack() returned nil")
	}

	if s.len != 0 {
		t.Errorf("Expected s.len to be 0, got %d", s.len)
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

func TestNewStackFromSlice(t *testing.T) {
	slice := []int{4, 3, 2, 1}
	s := NewStackFromSlice(slice)

	if s == nil {
		t.Error("NewStackFromSlice() returned nil")
	}

	if s.len != 4 {
		t.Errorf("Expected s.len to be 4, got %d", s.len)
	}

	if len(s.arr) != 4 {
		t.Errorf("Expected len(s.arr) to be 4, got %d", len(s.arr))
	}

	if val, _ := s.Top(); val != 1 {
		t.Errorf("Expected s.Top() to be 1, got %v", val)
	}

	expected := []int{1, 2, 3, 4}
	for i, exp := range expected {
		val, err := s.Pop()
		if err != nil {
			t.Errorf("At index %d got error: %v", i, err)
		}
		if !reflect.DeepEqual(val, exp) {
			t.Errorf("At index %d: expected %v, got %v", i, exp, val)
		}
	}

	if s.len != 0 {
		t.Errorf("Expected s.len to be 0 after all pops, got %d", s.len)
	}

	strslice := []string{"c", "b", "a"}
	strs := NewStackFromSlice(strslice)

	expectedStr := []string{"a", "b", "c"}
	for i, exp := range expectedStr {
		top, err := strs.Pop()
		if err != nil {
			t.Errorf("At index %d got error: %v", i, err)
		}
		if !reflect.DeepEqual(top, exp) {
			t.Errorf("At index %d: expected %v, got %v", i, exp, top)
		}
	}

	empty := NewStackFromSlice([]int{})
	if empty.len != 0 {
		t.Errorf("Expected len=0 for empty slice, got %d", empty.len)
	}
	if len(empty.arr) != 0 {
		t.Errorf("Expected arr length=0 for empty slice, got %d", len(empty.arr))
	}

	bigSlice := make([]int, 100000)
	for i := range bigSlice {
		bigSlice[i] = i
	}
	bigStack := NewStackFromSlice(bigSlice)
	if bigStack.len != 65535 {
		t.Errorf("Expected len=65535 (max), got %d", bigStack.len)
	}
	if len(bigStack.arr) != 65535 {
		t.Errorf("Expected arr length=65535, got %d", len(bigStack.arr))
	}
	if bigStack.arr[0] != 0 {
		t.Errorf("Expected first element 0, got %v", bigStack.arr[0])
	}
	if bigStack.arr[65534] != 65534 {
		t.Errorf("Expected last element 65534, got %v", bigStack.arr[65534])
	}
}

func TestPush(t *testing.T) {
	s := NewStack[int]()

	err := s.Push(42)
	if err != nil {
		t.Errorf("Push() unexpected error: %v", err)
	}
	if s.len != 1 {
		t.Errorf("Expected len=1, got %d", s.len)
	}
	if len(s.arr) != 1 {
		t.Errorf("Expected arr length=1, got %d", len(s.arr))
	}
	if s.arr[0] != 42 {
		t.Errorf("Expected arr[0]=42, got %v", s.arr[0])
	}

	err = s.Push(100)
	if err != nil {
		t.Errorf("Push() unexpected error: %v", err)
	}
	if s.len != 2 {
		t.Errorf("Expected len=2, got %d", s.len)
	}
	if len(s.arr) != 2 {
		t.Errorf("Expected arr length=2, got %d", len(s.arr))
	}
	if s.arr[1] != 100 {
		t.Errorf("Expected arr[1]=100, got %v", s.arr[1])
	}

	err = s.Push(200)
	if err != nil {
		t.Errorf("Push() unexpected error: %v", err)
	}
	if s.len != 3 {
		t.Errorf("Expected len=3, got %d", s.len)
	}
	if s.arr[2] != 200 {
		t.Errorf("Expected arr[2]=200, got %v", s.arr[2])
	}
}

func TestPushOverflow(t *testing.T) {
	s := NewStack[int]()

	for i := 0; i < 65535; i++ {
		err := s.Push(i)
		if err != nil {
			t.Errorf("Push(%d) unexpected error: %v", i, err)
		}
	}

	if s.len != 65535 {
		t.Errorf("Expected len=65535, got %d", s.len)
	}

	err := s.Push(65535)
	if err != errors.ErrorOverflow {
		t.Errorf("Expected ErrorOverflow, got %v", err)
	}
	if s.len != 65535 {
		t.Errorf("Expected len still 65535, got %d", s.len)
	}
}

func TestPop(t *testing.T) {
	s := NewStackFromSlice([]int{1, 2, 3, 4})

	expected := []int{4, 3, 2, 1}
	for i, exp := range expected {
		val, err := s.Pop()
		if err != nil {
			t.Errorf("At index %d got unexpected error: %v", i, err)
		}
		if val != exp {
			t.Errorf("At index %d: expected %d, got %d", i, exp, val)
		}
	}

	if s.len != 0 {
		t.Errorf("Expected len=0 after all pops, got %d", s.len)
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

func TestTop(t *testing.T) {
	s := NewStackFromSlice([]int{1, 2, 3})

	val, err := s.Top()
	if err != nil {
		t.Errorf("Top() unexpected error: %v", err)
	}
	if val != 3 {
		t.Errorf("Expected Top()=3, got %v", val)
	}

	if s.len != 3 {
		t.Errorf("Expected len=3 after Top(), got %d", s.len)
	}
	if len(s.arr) != 3 {
		t.Errorf("Expected arr length=3 after Top(), got %d", len(s.arr))
	}

	popped, _ := s.Pop()
	if popped != 3 {
		t.Errorf("Expected Pop()=3 after Top(), got %v", popped)
	}

	empty := NewStack[int]()
	val, err = empty.Top()
	if err != errors.ErrorEmpty {
		t.Errorf("Expected ErrorEmpty from empty stack, got %v", err)
	}
	var zero int
	if val != zero {
		t.Errorf("Expected zero value (0), got %v", val)
	}
}

func TestLen(t *testing.T) {
	s := NewStack[int]()

	if s.Len() != 0 {
		t.Errorf("Expected len=0, got %d", s.Len())
	}

	s.Push(1)
	if s.Len() != 1 {
		t.Errorf("Expected len=1, got %d", s.Len())
	}

	s.Push(2)
	s.Push(3)
	if s.Len() != 3 {
		t.Errorf("Expected len=3, got %d", s.Len())
	}

	s.Pop()
	if s.Len() != 2 {
		t.Errorf("Expected len=2 after Pop, got %d", s.Len())
	}

	s.Clear()
	if s.Len() != 0 {
		t.Errorf("Expected len=0 after Clear, got %d", s.Len())
	}
}

func TestIsEmpty(t *testing.T) {
	s := NewStack[int]()

	if !s.IsEmpty() {
		t.Error("Expected IsEmpty() to be true for new stack")
	}

	s.Push(1)
	if s.IsEmpty() {
		t.Error("Expected IsEmpty() to be false after Push")
	}

	s.Pop()
	if !s.IsEmpty() {
		t.Error("Expected IsEmpty() to be true after Pop")
	}

	s.Push(1)
	s.Push(2)
	s.Clear()
	if !s.IsEmpty() {
		t.Error("Expected IsEmpty() to be true after Clear")
	}
}

func TestClear(t *testing.T) {
	s := NewStackFromSlice([]int{1, 2, 3, 4, 5})

	if s.len != 5 {
		t.Errorf("Expected len=5 before Clear, got %d", s.len)
	}
	if len(s.arr) != 5 {
		t.Errorf("Expected arr length=5 before Clear, got %d", len(s.arr))
	}

	s.Clear()

	if s.len != 0 {
		t.Errorf("Expected len=0 after Clear, got %d", s.len)
	}
	if len(s.arr) != 0 {
		t.Errorf("Expected arr length=0 after Clear, got %d", len(s.arr))
	}
	if s.arr == nil {
		t.Error("Expected arr to be initialized (empty slice), got nil")
	}

	_, err := s.Pop()
	if err != errors.ErrorEmpty {
		t.Errorf("Expected ErrorEmpty after Clear, got %v", err)
	}

	_, err = s.Top()
	if err != errors.ErrorEmpty {
		t.Errorf("Expected ErrorEmpty after Clear, got %v", err)
	}
}

func TestSetTop(t *testing.T) {
	s := NewStackFromSlice([]int{1, 2, 3})

	err := s.SetTop(100)
	if err != nil {
		t.Errorf("SetTop() unexpected error: %v", err)
	}

	val, _ := s.Top()
	if val != 100 {
		t.Errorf("Expected Top()=100 after SetTop, got %v", val)
	}

	if s.len != 3 {
		t.Errorf("Expected len=3 after SetTop, got %d", s.len)
	}
	if len(s.arr) != 3 {
		t.Errorf("Expected arr length=3 after SetTop, got %d", len(s.arr))
	}

	expected := []int{1, 2, 100}
	result := s.ToSlice()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	empty := NewStack[int]()
	err = empty.SetTop(42)
	if err != errors.ErrorEmpty {
		t.Errorf("Expected ErrorEmpty, got %v", err)
	}
}

func TestSetTopWithDifferentTypes(t *testing.T) {
	s := NewStackFromSlice([]any{1, "hello", 3.14})

	err := s.SetTop("new value")
	if err != nil {
		t.Errorf("SetTop() unexpected error: %v", err)
	}

	val, _ := s.Top()
	if val != "new value" {
		t.Errorf("Expected Top()='new value', got %v", val)
	}

	expected := []any{1, "hello", "new value"}
	result := s.ToSlice()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestExists(t *testing.T) {
	s := NewStackFromSlice([]any{1, "hello", 3.14, []int{1, 2}})

	if !s.Exists("hello") {
		t.Error("Expected Exists('hello') to be true")
	}

	if !s.Exists(1) {
		t.Error("Expected Exists(1) to be true")
	}

	if !s.Exists(3.14) {
		t.Error("Expected Exists(3.14) to be true")
	}

	if !s.Exists([]int{1, 2}) {
		t.Error("Expected Exists([1 2]) to be true")
	}

	if s.Exists("nonexistent") {
		t.Error("Expected Exists('nonexistent') to be false")
	}

	if s.Exists(999) {
		t.Error("Expected Exists(999) to be false")
	}

	if s.Exists([]int{3, 4}) {
		t.Error("Expected Exists([3 4]) to be false")
	}

	empty := NewStack[int]()
	if empty.Exists(42) {
		t.Error("Expected Exists(42) to be false for empty stack")
	}
}

func TestToSlice(t *testing.T) {
	empty := NewStack[int]()
	result := empty.ToSlice()
	if len(result) != 0 {
		t.Errorf("Expected empty slice, got %v", result)
	}

	s := NewStackFromSlice([]int{1, 2, 3, 4})
	result = s.ToSlice()

	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	if s.len != 4 {
		t.Errorf("Expected len=4 after ToSlice, got %d", s.len)
	}
	if len(s.arr) != 4 {
		t.Errorf("Expected arr length=4 after ToSlice, got %d", len(s.arr))
	}

	s.Push(5)
	if len(result) != 4 {
		t.Errorf("Expected result length still 4, got %d", len(result))
	}
}

func TestConcurrentOperations(t *testing.T) {
	s := NewStack[int]()
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(val int) {
			s.Push(val)
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	if s.Len() != 10 {
		t.Errorf("Expected len=10, got %d", s.Len())
	}

	for i := 0; i < 10; i++ {
		go func() {
			s.Pop()
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	if s.Len() != 0 {
		t.Errorf("Expected len=0 after all pops, got %d", s.Len())
	}
}

func TestConcurrentPushOverflow(t *testing.T) {
	s := NewStack[int]()
	done := make(chan bool)

	for i := 0; i < 65535; i++ {
		s.Push(i)
	}

	for i := 0; i < 10; i++ {
		go func() {
			err := s.Push(999)
			if err != errors.ErrorOverflow {
				t.Errorf("Expected ErrorOverflow, got %v", err)
			}
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	if s.Len() != 65535 {
		t.Errorf("Expected len still 65535, got %d", s.Len())
	}
}

func BenchmarkPush(b *testing.B) {
	s := NewStack[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	s := NewStack[int]()
	for i := 0; i < 100000; i++ {
		s.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func BenchmarkTop(b *testing.B) {
	s := NewStack[int]()
	for i := 0; i < 100000; i++ {
		s.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Top()
	}
}

func BenchmarkSetTop(b *testing.B) {
	s := NewStack[int]()
	for i := 0; i < 100000; i++ {
		s.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.SetTop(i)
	}
}

func BenchmarkExists(b *testing.B) {
	s := NewStack[any]()
	for i := 0; i < 100000; i++ {
		s.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Exists(rand.Intn(65535))
	}
}

func BenchmarkToSlice(b *testing.B) {
	s := NewStack[int]()
	for i := 0; i < 100000; i++ {
		s.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.ToSlice()
	}
}
