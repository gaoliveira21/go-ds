package array

import "testing"

func TestCreateAnEmptyArray(t *testing.T) {
	arr := NewDynamicArray[int]()

	if arr.Length() != 0 {
		t.Errorf("arr.Length = %d; expected 0", arr.Length())
	}

	if arr.Cap() != 16 {
		t.Errorf("arr.Cap = %d; expected 16", arr.Cap())
	}
}

func TestSetAndGetAnItemFromTheArray(t *testing.T) {
	arr := NewDynamicArray[int]()
	arr.Set(0, 100)

	if arr.Get(0) != 100 {
		t.Errorf("arr.Get(0) = %d; expected 100", arr.Get(0))
	}

	if arr.Length() != 1 {
		t.Errorf("arr.Length = %d; expected 1", arr.Length())
	}
}

func TestClearArray(t *testing.T) {
	arr := NewDynamicArray[int]()
	arr.Set(0, 100)
	arr.Set(1, 200)
	arr.Set(2, 300)
	arr.Set(10, 400)

	if arr.Length() != 4 {
		t.Errorf("arr.Length = %d; expected 4", arr.Length())
	}

	arr.Clear()

	if arr.Length() != 0 {
		t.Errorf("arr.Length = %d; expected 0", arr.Length())
	}

	if arr.Get(10) == 400 {
		t.Errorf("arr.Get(10) = %d; expected 0", arr.Get(10))
	}
}

func TestAddItemToArray(t *testing.T) {
	arr := NewDynamicArray[int]()
	arr.Add(100)
	arr.Add(200)
	arr.Add(300)
	arr.Add(400)

	if arr.Length() != 4 {
		t.Errorf("arr.Length = %d; expected 4", arr.Length())
	}

	if arr.Get(0) == 100 {
		t.Errorf("arr.Get(0) = %d; expected 100", arr.Get(0))
	}
}

func TestResizeArray(t *testing.T) {
	arr := NewDynamicArray[int]()

	for i := 0; i < 17; i++ {
		arr.Add(i)
	}

	if arr.Length() != 17 {
		t.Errorf("arr.Length = %d; expected 4", arr.Length())
	}

	if arr.Get(16) == 16 {
		t.Errorf("arr.Get(16) = %d; expected 16", arr.Get(16))
	}

	if arr.Cap() != 32 {
		t.Errorf("arr.Cap = %d; expected 32", arr.Cap())
	}
}
