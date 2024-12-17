package main

import (
	"reflect"
	"testing"
)

func checkStackLength[T Ordered](t *testing.T, stack Stack[T], length int, comp func(i, j int) bool) bool {
	if comp != nil && !comp(stack.Len(), length) {
		t.Error("stack doesn't match required condition")
		return false
	} else if comp == nil && stack.Len() != length {
		t.Errorf("empty stack should have a length of %d; stack has length of %d", length, stack.Len())
		return false
	}
	return true
}

func checkStackValues[T Ordered](t *testing.T, stack Stack[T], expectedValues []T) bool {
	if !reflect.DeepEqual(stack.values, expectedValues) {
		t.Errorf("stack doesn't match expected values; has values: %v, expected: %v", stack.values, expectedValues)
		return false
	}
	return true
}

func checkStackTop[T Ordered](t *testing.T, stack Stack[T], expectedTop T) bool {
	var validLength = func(stackLength, expectedLength int) bool {
		return stackLength > expectedLength
	}

	if !checkStackLength(t, stack, 0, validLength) {
		return false
	}

	n := stack.Len() - 1
	if stack.values[n] != expectedTop {
		t.Errorf("stack top doesn't match expected value: expected %v, got %v", expectedTop, stack.values[n])
		return false
	}
	return true
}

func TestStackPush(t *testing.T) {
	s := Stack[int]{}
	checkStackLength(t, s, 0, nil)

	s.Push(1)
	checkStackLength(t, s, 1, nil)
	checkStackValues(t, s, []int{1})

	s.Push(10)
	checkStackLength(t, s, 2, nil)
	checkStackValues(t, s, []int{1, 10})
}

func TestStackPop(t *testing.T) {
	s := Stack[int]{[]int{1, 2, 3, 4, 5, 6}}

	last := s.values[s.Len()-1]
	top := s.Pop()
	if *top != last {
		t.Errorf("popped and got the wrong values; expected %d, got %d", last, top)
	}
	checkStackTop(t, s, 5)
}

func TestStackPeek(t *testing.T) {
	s := Stack[int]{[]int{1, 2, 3, 4, 5, 6}}
	top := s.Peek()

	checkStackTop(t, s, *top)
	s.Push(5)
	checkStackTop(t, s, *s.Peek())
}
