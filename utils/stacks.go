package utils

import (
	"fmt"
)

type List[T any] struct {
	vals []T
}

func ListFrom[T any](initial []T) List[T] {
	return List[T]{vals: append([]T{}, initial...)}
}

func (l *List[T]) Length() int {
	return len(l.vals)
}

func (l *List[T]) Get(i int) (T, error) {
	var zero T
	if len(l.vals) <= i {
		return zero, fmt.Errorf("index %d not in range", i)
	}
	return l.vals[i], nil
}

func (s *List[T]) Push(val ...T) {
	s.vals = append(s.vals, val...)
}

func (s *List[T]) Queue(val ...T) {
	s.vals = append(s.vals, val...)
}

func (l *List[T]) Dequeue() T {
	bot := l.vals[0]
	var new []T
	if len(l.vals) > 1 {
		new = l.vals[1:]
	}
	l.vals = new
	return bot
}

func (s *List[T]) Pop() T {
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top
}

func (s *List[T]) PopN(n int) []T {
	topN := s.vals[len(s.vals)-n:]
	s.vals = s.vals[:len(s.vals)-n]
	return topN
}

func (s *List[T]) Count() int {
	return len(s.vals)
}

func (s *List[T]) Slice() []T {
	return append([]T{}, s.vals...)
}