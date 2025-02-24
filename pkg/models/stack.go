package models

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(data T) {
	s.items = append(s.items, data)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}
	result := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return result
}

func (s *Stack[T]) Top() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}
	return s.items[len(s.items)-1]
}
