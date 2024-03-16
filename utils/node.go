package linkedlist

type node[T comparable] struct {
	Value T
	next  *node[T]
}

func (n node[T]) Offset(i uint32) (node[T], bool) {
	if i > 0 {
		if n.next == nil {
			return node[T]{}, false
		}
		return n.next.Offset(i - 1)
	}
	return n, true
}

func (n *node[T]) removeFirstValueMatch(v T) bool {
	if n != nil && n.Value == v {
		if n.next != nil {
			n.Value = n.next.Value
			n.next = n.next.next
		}
		return true
	} else if n == nil || n.next == nil {
		return false
	}
	return n.next.removeFirstValueMatch(v)
}
