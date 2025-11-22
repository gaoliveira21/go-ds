package singly_linked_list

import "fmt"

type LinkedList[T any] struct {
	head   *LinkedListNode[T]
	tail   *LinkedListNode[T]
	Length int
}

func NewList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (list *LinkedList[T]) GetFirst() (T, bool) {
	if list.head == nil {
		return *new(T), false
	}

	return list.head.value, true
}

func (list *LinkedList[T]) DeleteFirst() {
	if list.head != nil {
		list.head = list.head.next
		list.Length--
	}
}

func (list *LinkedList[T]) Append(v T) {
	node := &LinkedListNode[T]{
		value: v,
	}

	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		list.tail = node
	}

	list.Length += 1
}

func (list *LinkedList[T]) Prepend(v T) {
	node := &LinkedListNode[T]{
		value: v,
	}

	node.next = list.head
	list.head = node

	if list.tail == nil {
		list.tail = node
	}

	list.Length += 1
}

func (list *LinkedList[T]) Contains(fn func(v T) bool) (bool, T) {
	if list.head == nil {
		return false, *new(T)
	}

	n := list.head

	for n != nil {
		if fn(n.value) {
			return true, n.value
		}

		n = n.next
	}

	return false, *new(T)
}

func (list *LinkedList[T]) Delete(fn func(v T) bool) {
	if list.head == nil {
		return
	}

	if list.Length == 1 && fn(list.head.value) {
		list.head = nil
		list.tail = nil
		list.Length--
		return
	}

	if fn(list.head.value) {
		list.head = list.head.next
		list.Length--
		return
	}

	n := list.head

	for n.next != nil {
		if fn(n.next.value) {
			n.next = n.next.next
			list.Length--
			break
		}

		n = n.next
	}

	if fn(list.tail.value) {
		list.tail = n
	}
}

func (list *LinkedList[T]) Join(separator string) string {
	n := list.head

	result := ""

	for n != nil {
		result += n.String()

		if n.next != nil && list.Length > 1 {
			result += separator
		}

		n = n.next
	}

	fmt.Printf("%s", result)

	return result
}

func (list *LinkedList[T]) Reverse() {
	var current *LinkedListNode[T] = list.head
	var prev *LinkedListNode[T]
	var next *LinkedListNode[T]

	for current != nil {
		next = current.next
		current.next = prev

		prev = current
		current = next
	}

	list.tail = list.head
	list.head = prev
}
