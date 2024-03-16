package linkedlist

type node[T comparable] struct {
	value T
	next  *node[T]
}

func (n *node[T]) stepForward(i uint32) *node[T] {
	if i > 0 {
		return n.next.stepForward(i - 1)
	}
	return n
}

func (n *node[T]) removeFirstValueMatch(v T) bool {
	if n.value == v {
		if n.next == nil {
			n = nil
		} else {
			n.value = n.next.value
			n.next = n.next.next
		}
		return true
	} else if n.next == nil {
		return false
	}
	return n.next.removeFirstValueMatch(v)
}
