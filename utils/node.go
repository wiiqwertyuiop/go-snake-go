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

func (n *node[T]) removeFirstValueMatchAfterHead(v T) bool {
	if n.next == nil {
		return false
	} else if n.next.Value == v {
		if n.next.next == nil {
			n.next = nil
		} else {
			n.next = n.next.next
		}
		return true
	}
	return n.next.removeFirstValueMatchAfterHead(v)
}

func (n *node[T]) popOnOffsetAfterHead(i uint32) (T, bool) {
	if n.next == nil {
		return node[T]{}.Value, false
	} else if i > 0 {
		return n.next.popOnOffsetAfterHead(i - 1)
	}
	val := n.next.Value
	if n.next.next == nil {
		n.next = nil
	} else {
		n.next = n.next.next
	}
	return val, true
}
