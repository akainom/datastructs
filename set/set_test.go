package set

import (
	"datastructs/errors"
	"reflect"
	"sort"
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewSet[int]()

	if s == nil {
		t.Error("NewSet() returned nil")
	}

	if s.items == nil {
		t.Error("Expected s.items to be initialized, got nil")
	}

	if len(s.items) != 0 {
		t.Errorf("Expected len(s.items) to be 0, got %d", len(s.items))
	}

	if s.m == nil {
		t.Error("Expected s.m to be initialized, got nil")
	}
}

func TestNewSetFromSlice(t *testing.T) {
	empty := NewSetFromSlice([]int{})
	if len(empty.items) != 0 {
		t.Errorf("Expected len=0, got %d", len(empty.items))
	}

	slice := []int{1, 2, 3, 2, 1, 4, 5, 3}
	s := NewSetFromSlice(slice)

	expected := []int{1, 2, 3, 4, 5}
	result := s.ToSlice()
	sort.Ints(result)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	if s.Len() != 5 {
		t.Errorf("Expected len=5, got %d", s.Len())
	}

	strSlice := []string{"a", "b", "a", "c", "b", "d"}
	strSet := NewSetFromSlice(strSlice)
	expectedStr := []string{"a", "b", "c", "d"}
	resultStr := strSet.ToSlice()
	sort.Strings(resultStr)

	if !reflect.DeepEqual(resultStr, expectedStr) {
		t.Errorf("Expected %v, got %v", expectedStr, resultStr)
	}

	bigSlice := make([]int, 100000000)
	for i := range bigSlice {
		bigSlice[i] = i
	}
	bigSet := NewSetFromSlice(bigSlice)
	if len(bigSet.items) != 17359360 {
		t.Errorf("Expected arr length=65535 (max), got %d", len(bigSet.items))
	}
}

func TestAdd(t *testing.T) {
	s := NewSet[int]()

	err := s.Add(42)
	if err != nil {
		t.Errorf("Add() unexpected error: %v", err)
	}
	if s.Len() != 1 {
		t.Errorf("Expected len=1, got %d", s.Len())
	}
	if !s.Exists(42) {
		t.Error("Expected 42 to exist in set")
	}

	err = s.Add(42)
	if err != nil {
		t.Errorf("Add(42) unexpected error: %v", err)
	}
	if s.Len() != 1 {
		t.Errorf("Expected len still 1, got %d", s.Len())
	}

	s2 := NewSet[string]()
	s2.Add("hello")
	if s2.Len() != 1 {
		t.Errorf("Expected len=1, got %d", s2.Len())
	}
	if !s2.Exists("hello") {
		t.Error("Expected 'hello' to exist in set")
	}
}

func TestAddOverflow(t *testing.T) {
	s := NewSet[int]()

	for i := 0; i < 17359360; i++ {
		err := s.Add(i)
		if err != nil {
			t.Errorf("Add(%d) unexpected error: %v", i, err)
		}
	}

	if s.Len() != 17359360 {
		t.Errorf("Expected len=17359360, got %d", s.Len())
	}

	err := s.Add(999)
	if err != errors.ErrorOverflow {
		t.Errorf("Expected ErrorOverflow, got %v", err)
	}
}

func TestExists(t *testing.T) {
	s := NewSetFromSlice([]int{1, 2, 3, 4, 5})

	if !s.Exists(1) {
		t.Error("Expected 1 to exist")
	}
	if !s.Exists(3) {
		t.Error("Expected 3 to exist")
	}
	if !s.Exists(5) {
		t.Error("Expected 5 to exist")
	}

	if s.Exists(0) {
		t.Error("Expected 0 to not exist")
	}
	if s.Exists(6) {
		t.Error("Expected 6 to not exist")
	}
	if s.Exists(100) {
		t.Error("Expected 100 to not exist")
	}

	empty := NewSet[int]()
	if empty.Exists(42) {
		t.Error("Expected 42 to not exist in empty set")
	}
}

func TestRemove(t *testing.T) {
	s := NewSetFromSlice([]int{1, 2, 3, 4, 5})

	err := s.Remove(3)
	if err != nil {
		t.Errorf("Remove(3) unexpected error: %v", err)
	}
	if s.Len() != 4 {
		t.Errorf("Expected len=4, got %d", s.Len())
	}
	if s.Exists(3) {
		t.Error("Expected 3 to be removed")
	}

	err = s.Remove(999)
	if err != errors.ErrorNF {
		t.Errorf("Expected ErrorNF, got %v", err)
	}
	if s.Len() != 4 {
		t.Errorf("Expected len still 4, got %d", s.Len())
	}

	empty := NewSet[int]()
	err = empty.Remove(42)
	if err != errors.ErrorNF {
		t.Errorf("Expected ErrorNF from empty set, got %v", err)
	}
}

func TestLen(t *testing.T) {
	s := NewSet[int]()

	if s.Len() != 0 {
		t.Errorf("Expected len=0, got %d", s.Len())
	}

	s.Add(1)
	if s.Len() != 1 {
		t.Errorf("Expected len=1, got %d", s.Len())
	}

	s.Add(2)
	s.Add(3)
	if s.Len() != 3 {
		t.Errorf("Expected len=3, got %d", s.Len())
	}

	s.Remove(2)
	if s.Len() != 2 {
		t.Errorf("Expected len=2 after Remove, got %d", s.Len())
	}

	s.Clear()
	if s.Len() != 0 {
		t.Errorf("Expected len=0 after Clear, got %d", s.Len())
	}
}

func TestIsEmpty(t *testing.T) {
	s := NewSet[int]()

	if !s.IsEmpty() {
		t.Error("Expected IsEmpty() to be true for new set")
	}

	s.Add(1)
	if s.IsEmpty() {
		t.Error("Expected IsEmpty() to be false after Add")
	}

	s.Remove(1)
	if !s.IsEmpty() {
		t.Error("Expected IsEmpty() to be true after Remove")
	}

	s.Add(1)
	s.Add(2)
	s.Clear()
	if !s.IsEmpty() {
		t.Error("Expected IsEmpty() to be true after Clear")
	}
}

func TestClear(t *testing.T) {
	s := NewSetFromSlice([]int{1, 2, 3, 4, 5})

	if s.Len() != 5 {
		t.Errorf("Expected len=5 before Clear, got %d", s.Len())
	}

	s.Clear()

	if s.Len() != 0 {
		t.Errorf("Expected len=0 after Clear, got %d", s.Len())
	}
	if s.items == nil {
		t.Error("Expected items to be initialized (empty map), got nil")
	}
	if len(s.items) != 0 {
		t.Errorf("Expected items length=0, got %d", len(s.items))
	}

	if s.Exists(1) {
		t.Error("Expected 1 to be removed after Clear")
	}
	if s.Exists(5) {
		t.Error("Expected 5 to be removed after Clear")
	}
}

func TestToSlice(t *testing.T) {
	empty := NewSet[int]()
	result := empty.ToSlice()
	if len(result) != 0 {
		t.Errorf("Expected empty slice, got %v", result)
	}

	s := NewSetFromSlice([]int{1, 2, 3, 4, 5})
	result = s.ToSlice()

	expected := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
	if len(result) != 5 {
		t.Errorf("Expected len=5, got %d", len(result))
	}
	for _, v := range result {
		if !expected[v] {
			t.Errorf("Unexpected value: %d", v)
		}
		delete(expected, v)
	}
	if len(expected) != 0 {
		t.Errorf("Missing values: %v", expected)
	}

	if s.Len() != 5 {
		t.Errorf("Expected len still 5, got %d", s.Len())
	}
}

func TestConcurrentOperations(t *testing.T) {
	s := NewSet[int]()
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(val int) {
			s.Add(val)
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
		go func(val int) {
			if !s.Exists(val) {
				t.Errorf("Expected %d to exist", val)
			}
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	for i := 0; i < 10; i++ {
		go func(val int) {
			s.Remove(val)
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	if s.Len() != 0 {
		t.Errorf("Expected len=0 after all removes, got %d", s.Len())
	}
}

func TestConcurrentPushOverflow(t *testing.T) {
	s := NewSet[int]()
	done := make(chan bool)

	for i := 0; i < 17359360; i++ {
		err := s.Add(i)
		if err != nil {
			t.Errorf("Add(%d) unexpected error: %v", i, err)
		}
	}

	for i := 0; i < 10; i++ {
		go func() {
			err := s.Add(999)
			if err != errors.ErrorOverflow {
				t.Errorf("Expected ErrorOverflow, got %v", err)
			}
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	if s.Len() != 17359360 {
		t.Errorf("Expected len still 17359360, got %d", s.Len())
	}
}

func BenchmarkAdd(b *testing.B) {
	s := NewSet[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
}

func BenchmarkExists(b *testing.B) {
	s := NewSet[int]()
	for i := 0; i < 100000; i++ {
		s.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Exists(i % 100000)
	}
}

func BenchmarkRemove(b *testing.B) {
	s := NewSet[int]()
	for i := 0; i < 100000; i++ {
		s.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Remove(i % 100000)
	}
}

func BenchmarkToSlice(b *testing.B) {
	s := NewSet[int]()
	for i := 0; i < 100000; i++ {
		s.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.ToSlice()
	}
}
