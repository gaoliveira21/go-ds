package stack

import (
	"testing"
)

func TestIsStackEmpty(t *testing.T) {
	s := NewStack[int]()

	if s.IsEmpty() == false {
		t.Error("stack should be empty; expected = true")
	}
}

func TestIsStackNotEmpty(t *testing.T) {
	s := NewStack[int]()

	s.Push(1)

	if s.IsEmpty() == true {
		t.Error("stack should not be empty; expected = false")
	}
}

func TestPopValueFromStack(t *testing.T) {
	s := NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	v, _ := s.Pop()

	if v != 3 {
		t.Errorf("invalid value returned from stack; expected = 3; received %d", v)
	}
}

func TestPopFromEmptyStack(t *testing.T) {
	s := NewStack[int]()

	_, found := s.Pop()

	if found != false {
		t.Errorf("expected = false; received true")
	}
}
