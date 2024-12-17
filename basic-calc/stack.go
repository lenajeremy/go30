package main

type Ordered interface {
	~int | ~uint | ~float64 | ~float32 | ~uint8 | ~uint16 | ~uint32
}

type Stack[T Ordered] struct {
	values []T
}

func (s *Stack[T]) Len() int {
	return len(s.values)
}

func (s *Stack[T]) Pop() *T {
	n := len(s.values)
	if n == 0 {
		return nil
	}
	res := s.values[n-1]
	s.values = s.values[:n-1]
	return &res
}

func (s *Stack[T]) Push(v T) {
	s.values = append(s.values, v)
}

func (s *Stack[T]) Peek() *T {
	n := len(s.values)
	if n == 0 {
		return nil
	}
	return &s.values[n-1]
}
