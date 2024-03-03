package queue

import "github.com/gaoliveira21/go-ds/linked_list"

type Queue[T any] struct {
	list *linked_list.LinkedList[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		list: linked_list.NewList[T](),
	}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.Length == 0
}

func (q *Queue[T]) Enqueue(v T) {
	q.list.Append(v)
}

func (q *Queue[T]) Peek() (T, bool) {
	return q.list.GetFirst()
}

func (q *Queue[T]) Dequeue() (T, bool) {
	v, f := q.list.GetFirst()

	q.list.DeleteFirst()

	return v, f
}
