// Small wrapper class around node's to keep track of the head and size of list
// Also provides some safety
//

package linkedlist

type LinkedList[T comparable] struct {
	head *node[T]
	size uint32
}

func (list *LinkedList[T]) Size() uint32 {
	return list.size
}

func (list *LinkedList[T]) Prepend(val T) *LinkedList[T] {
	list.head = &node[T]{Value: val, next: list.head}
	list.size++
	return list
}

func (list *LinkedList[T]) RemoveFirstValueMatch(value T) *LinkedList[T] {
	n := node[T]{next: list.head}
	if n.removeFirstValueMatchAfterHead(value) {
		list.size--
		list.head = n.next
	}
	return list
}

func (list *LinkedList[T]) ReadNodeOnIndex(index uint32) (node[T], bool) {
	if list.head == nil {
		return node[T]{}, false
	}
	return list.head.Offset(index)
}

func (list *LinkedList[T]) ReadValueOnIndex(index uint32) (T, bool) {
	if list.head == nil {
		return node[T]{}.Value, false
	}
	n, found := list.head.Offset(index)
	return n.Value, found
}

func (list *LinkedList[T]) PopWithValueOnIndex(index uint32) (T, bool) {
	n := node[T]{next: list.head}
	val, found := n.popOnOffsetAfterHead(index)
	if found {
		list.size--
		list.head = n.next
	}
	return val, found
}
