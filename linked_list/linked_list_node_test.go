package linked_list

import (
	"testing"
)

func TestCreateNodeWithValue(t *testing.T) {
	node := LinkedListNode[int]{
		value: 10,
	}

	if node.value != 10 {
		t.Errorf("node.value = %d; expected %d", node.value, 10)
	}

	if node.next != nil {
		t.Errorf("node.next = %s; expected nil", node.next)
	}
}

func TestCastingNodeValueToString(t *testing.T) {
	node := LinkedListNode[int]{
		value: 10,
	}

	str := node.String()

	if str != "10" {
		t.Errorf("str = %s; expected '10'", str)
	}
}
