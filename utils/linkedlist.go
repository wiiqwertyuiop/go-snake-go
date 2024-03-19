// Customized linked list for snake!
//
// Small wrapper class around node's to keep track of the head and size of list
// Also provides some helper functions and safety

package linkedlist

type LinkedList[T comparable] struct {
	head Node[T]
	size int
}

func (list *LinkedList[T]) Size() int {
	return list.size
}

func (list *LinkedList[T]) Prepend(val T) *LinkedList[T] {
	list.head.next = &Node[T]{Value: val, next: list.head.next}
	list.size++
	return list
}

func (list *LinkedList[T]) RemoveFirstValueMatch(value T) *LinkedList[T] {
	if list.head.removeFirstValueMatchAfterHead(value) {
		list.size--
	}
	return list
}

func (list *LinkedList[T]) GetNodeOnIndex(index int) (*Node[T], bool) {
	return list.head.Offset(index + 1)
}

func (list *LinkedList[T]) ReadValueOnIndex(index int) (T, bool) {
	n, found := list.head.Offset(index + 1)
	return n.Value, found
}

func (list *LinkedList[T]) PopWithValueOnIndex(index int) (T, bool) {
	val, found := list.head.popOnOffsetAfterHead(index)
	if found {
		list.size--
	}
	return val, found
}

func (list *LinkedList[T]) LoopActionOnListRemoveNextOnFalse(callback func(curNode *Node[T]) bool) {
	for n := list.head.next; n != nil; n = n.next {
		if !callback(n) {
			// NOTE: i only use this to chop off the bit of the snake...
			// Ideally you would want to make sure there is nothing else connected after it
			n.next = nil
			list.size--
			break
		}
	}
}
