package linked_list

import (
	"testing"
)

func TestCreateEmptyList(t *testing.T) {
	list := NewList[int]()

	if list.Length != 0 {
		t.Errorf("list.Length = %d; expected 0", list.Length)
	}

	if list.head != nil {
		t.Errorf("list.head expected to be nil")
	}

	if list.tail != nil {
		t.Errorf("list.head expected to be nil")
	}
}

func TestAppendToEmptyList(t *testing.T) {
	list := NewList[int]()

	list.Append(1)

	if list.Length != 1 {
		t.Errorf("list.Length = %d; expected 1", list.Length)
	}

	if list.head.value != 1 {
		t.Errorf("list.head = %d; expected 1", list.head.value)
	}

	if list.tail.value != 1 {
		t.Errorf("list.tail = %d; expected 1", list.tail.value)
	}
}

func TestAppendToList(t *testing.T) {
	list := NewList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	if list.Length != 3 {
		t.Errorf("list.Length = %d; expected 3", list.Length)
	}

	if list.head.value != 1 {
		t.Errorf("list.head = %d; expected 1", list.head.value)
	}

	if list.tail.value != 3 {
		t.Errorf("list.tail = %d; expected 3", list.tail.value)
	}
}

func TestPrependToEmptyList(t *testing.T) {
	list := NewList[int]()

	list.Prepend(1)

	if list.Length != 1 {
		t.Errorf("list.Length = %d; expected 1", list.Length)
	}

	if list.head.value != 1 {
		t.Errorf("list.head = %d; expected 1", list.head.value)
	}

	if list.tail.value != 1 {
		t.Errorf("list.tail = %d; expected 1", list.tail.value)
	}
}

func TestPrependToList(t *testing.T) {
	list := NewList[int]()

	list.Append(2)
	list.Append(3)
	list.Prepend(1)

	if list.Length != 3 {
		t.Errorf("list.Length = %d; expected 3", list.Length)
	}

	if list.head.value != 1 {
		t.Errorf("list.head = %d; expected 1", list.head.value)
	}

	if list.tail.value != 3 {
		t.Errorf("list.tail = %d; expected 3", list.tail.value)
	}
}

func TestContainsInEmptyList(t *testing.T) {
	list := NewList[int]()

	found, v := list.Contains(func(v int) bool {
		return v == 5
	})

	if found != false || v != 0 {
		t.Errorf("Return values should be false and nil")
	}
}

func TestContainsInListWhenValueExists(t *testing.T) {
	list := NewList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(56)

	found, v := list.Contains(func(v int) bool {
		return v == 5
	})

	if found != true || v != 5 {
		t.Errorf("Value should be found in list")
	}
}

func TestContainsInListWhenValueNotExists(t *testing.T) {
	list := NewList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(56)

	found, v := list.Contains(func(v int) bool {
		return v == 100
	})

	if found != false || v != 0 {
		t.Errorf("Value should be not found in list")
	}
}

func TestDeleteHead(t *testing.T) {
	list := NewList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Delete(func(v int) bool {
		return v == 1
	})

	if list.head.value != 2 {
		t.Errorf("list.head.value = %d, expected = 2", list.head.value)
	}

	if list.Length != 2 {
		t.Errorf("list.Length = %d, expected = 2", list.Length)
	}
}

func TestDeleteNode(t *testing.T) {
	list := NewList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	list.Delete(func(v int) bool {
		return v == 3
	})

	found, _ := list.Contains(func(v int) bool {
		return v == 3
	})

	if found != false {
		t.Errorf("value 3 found in list")
	}

	if list.Length != 4 {
		t.Errorf("list.Length = %d, expected = 4", list.Length)
	}

	if list.head.next.next.value != 4 {
		t.Errorf("The third node value should be 4")
	}
}

func TestDeleteTail(t *testing.T) {
	list := NewList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	list.Delete(func(v int) bool {
		return v == 5
	})

	if list.Length != 4 {
		t.Errorf("list.Length = %d, expected = 4", list.Length)
	}

	if list.tail.value != 4 {
		t.Errorf("list.tail.value = %d, expected = 4", list.tail.value)
	}
}

func TestDeleteWhenThereIsOnlyOneElementInList(t *testing.T) {
	list := NewList[int]()

	list.Append(1)

	list.Delete(func(v int) bool {
		return v == 1
	})

	if list.head != nil {
		t.Errorf("list.head = %s, expected = nil", list.head)
	}

	if list.tail != nil {
		t.Errorf("list.tail = %s, expected = nil", list.tail)
	}

	if list.Length != 0 {
		t.Errorf("list.Length = %d, expected = nil", list.Length)
	}
}

func TestDeleteWithMultipleValueOccurrence(t *testing.T) {
	list := NewList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(2)
	list.Append(4)
	list.Append(5)

	list.Delete(func(v int) bool {
		return v == 2
	})

	if list.Length != 4 {
		t.Errorf("list.Length = %d, expected = 4", list.Length)
	}

	if list.head.next.next.value != 4 {
		t.Errorf("The third node value should be 4")
	}
}

func TestDeleteWhenListIsEmpty(t *testing.T) {
	list := NewList[int]()

	list.Delete(func(v int) bool {
		return v == 1
	})

	if list.head != nil {
		t.Errorf("list.head = %s, expected = nil", list.head)
	}

	if list.tail != nil {
		t.Errorf("list.tail = %s, expected = nil", list.tail)
	}

	if list.Length != 0 {
		t.Errorf("list.Length = %d, expected = nil", list.Length)
	}
}

func TestJoinList(t *testing.T) {
	list := NewList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	r := list.Join(",")

	if r != "1,2,3,4,5" {
		t.Errorf("result = %s, expected 1,2,3,4,5", r)
	}

	r = list.Join(" ")

	if r != "1 2 3 4 5" {
		t.Errorf("result = %s, expected 1 2 3 4 5", r)
	}
}

func TestJoinListWithSingleValue(t *testing.T) {
	list := NewList[int]()

	list.Append(1)

	r := list.Join(",")

	if r != "1" {
		t.Errorf("result = %s, expected 1", r)
	}
}

func TestReverseList(t *testing.T) {
	list := NewList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	list.Reverse()

	r := list.Join(" ")

	if r != "5 4 3 2 1" {
		t.Errorf("result = %s, expected 5 4 3 2 1", r)
	}

	if list.head.value != 5 {
		t.Errorf("list.head.value = %d, 5", list.head.value)
	}

	if list.tail.value != 1 {
		t.Errorf("list.tail.value = %d, 1", list.tail.value)
	}
}
