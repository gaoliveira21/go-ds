package queue

import (
	"testing"
)

func TestIsQueueEmpty(t *testing.T) {
	q := NewQueue[int]()

	if q.IsEmpty() == false {
		t.Error("queue should be empty; expected = true")
	}
}

func TestIsQueueNotEmpty(t *testing.T) {
	q := NewQueue[int]()

	q.Enqueue(1)

	if q.IsEmpty() == true {
		t.Error("queue should not be empty; expected = false")
	}
}

func TestPeekFirstValueFromQueue(t *testing.T) {
	q := NewQueue[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	v, _ := q.Peek()

	if v != 1 {
		t.Errorf("invalid value returned from queue; expected = 1; received %d", v)
	}
}

func TestPeekFirstValueFromEmptyQueue(t *testing.T) {
	q := NewQueue[int]()

	_, found := q.Peek()

	if found != false {
		t.Errorf("expected = false; received true")
	}
}

func TestDequeueValueFromQueue(t *testing.T) {
	q := NewQueue[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	v, _ := q.Dequeue()
	if v != 1 {
		t.Errorf("invalid value return from queue; expected = 1")
	}

	v, _ = q.Dequeue()
	if v != 2 {
		t.Errorf("invalid value return from queue; expected = 2")
	}

	v, _ = q.Dequeue()
	if v != 3 {
		t.Errorf("invalid value return from queue; expected = 3")
	}

	if q.IsEmpty() != true {
		t.Errorf("queue should be empty")
	}
}

func TestDequeueValueFromEmptyQueue(t *testing.T) {
	q := NewQueue[int]()

	v, f := q.Dequeue()
	if v != 0 || f != false {
		t.Errorf("invalid value return from queue; expected = 0")
	}
}
