package stack

import "github.com/gaoliveira21/go-ds/singly_linked_list"

type Stack[T any] struct {
	list *singly_linked_list.LinkedList[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		list: singly_linked_list.NewList[T](),
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.list.Length == 0
}

func (s *Stack[T]) Push(v T) {
	s.list.Prepend(v)
}

func (s *Stack[T]) Pop() (T, bool) {
	v, f := s.list.GetFirst()
	s.list.DeleteFirst()

	return v, f
}
