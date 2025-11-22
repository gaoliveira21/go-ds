package doubly_linked_list

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

func (list *LinkedList[T]) Append(v T) {
	node := &LinkedListNode[T]{
		value: v,
	}

	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		node.prev = list.tail
		list.tail = node
	}

	list.Length += 1
}

func (list *LinkedList[T]) Prepend(v T) {
	node := &LinkedListNode[T]{
		value: v,
	}

	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		node.next = list.head
		list.head.prev = node
		list.head = node
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
		list.head.prev = nil
		list.Length--
	}

	n := list.head

	for n != nil {
		if fn(n.value) {
			prevNode := n.prev
			nextNode := n.next

			prevNode.next = nextNode

			if list.tail == n {
				list.tail = prevNode
			} else {
				nextNode.prev = prevNode
			}

			list.Length--
			break
		}

		n = n.next
	}

	if fn(list.tail.value) {
		list.tail = list.tail.prev
		list.tail.next = nil
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

	return result
}

func (list *LinkedList[T]) Reverse() {
	n := list.head

	list.head, list.tail = list.tail, list.head

	for n != nil {
		n.prev, n.next = n.next, n.prev
		n = n.prev
	}
}
