package singly_linked_list

import "fmt"

type LinkedListNode[T any] struct {
	value T
	next  *LinkedListNode[T]
}

func (node *LinkedListNode[T]) String() string {
	return fmt.Sprint(node.value)
}
