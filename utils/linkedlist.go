package linkedlist

type LinkedList[T comparable] struct {
	head *node[T]
	size uint32
}

func (list *LinkedList[T]) Size() uint32 {
	return list.size
}

func (list *LinkedList[T]) GetValueByIndex(index uint32) (T, bool) {
	if list.size == 0 || list.size-1 < index {
		return node[T]{}.value, false
	}
	return list.head.stepForward(index).value, true
}

func (list *LinkedList[T]) Prepend(val T) *LinkedList[T] {
	list.head = &node[T]{value: val, next: list.head}
	list.size++
	return list
}

func (list *LinkedList[T]) RemoveFirstValueMatch(value T) *LinkedList[T] {
	if list.size != 0 && list.head.removeFirstValueMatch(value) {
		list.size--
	}
	return list
}
