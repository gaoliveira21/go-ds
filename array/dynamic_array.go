package array

type DynamicArray[T any] struct {
	len  uint
	cap  uint
	data []T
}

func NewDynamicArray[T any]() *DynamicArray[T] {
	return &DynamicArray[T]{
		len:  0,
		cap:  16,
		data: make([]T, 16),
	}
}

func (arr *DynamicArray[T]) Length() uint {
	return arr.len
}

func (arr *DynamicArray[T]) Cap() uint {
	return arr.cap
}

func (arr *DynamicArray[T]) Set(index int, el T) {
	arr.data[index] = el
	arr.len++
}

func (arr *DynamicArray[T]) Get(index int) T {
	return arr.data[index]
}

func (arr *DynamicArray[T]) Clear() {
	arr.data = make([]T, arr.cap)
	arr.len = 0
}

func (arr *DynamicArray[T]) Add(el T) {
	mustResize := arr.len+1 >= arr.cap

	if mustResize {
		arr.cap *= 2
		newArr := make([]T, arr.cap)
		copy(newArr, arr.data)
		arr.data = newArr
	}

	arr.len++
	arr.data[arr.len] = el
}
