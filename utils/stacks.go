package utils

type Stack[T any] struct {
	vals []T
}

func StackFrom[T any](initial []T) Stack[T] {
	return Stack[T]{vals: append([]T{}, initial...)}
}

func (s *Stack[T]) Push(val ...T) {
	s.vals = append(s.vals, val...)
}

func (s *Stack[T]) Pop() T {
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top
}

func (s *Stack[T]) PopN(n int) []T {
	topN := s.vals[len(s.vals)-n:]
	s.vals = s.vals[:len(s.vals)-n]
	return topN
}

func (s *Stack[T]) Count() int {
	return len(s.vals)
}

func (s *Stack[T]) Slice() []T {
	return append([]T{}, s.vals...)
}