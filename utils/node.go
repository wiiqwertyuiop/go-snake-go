package linkedlist

type Node[T comparable] struct {
	Value T
	next  *Node[T]
}

func (n Node[T]) Offset(i int) (*Node[T], bool) {
	if i > 0 {
		if n.next == nil {
			return &Node[T]{}, false
		}
		return n.next.Offset(i - 1)
	}
	return &n, true
}

func (n *Node[T]) removeFirstValueMatchAfterHead(v T) bool {
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

func (n *Node[T]) popOnOffsetAfterHead(i int) (T, bool) {
	if n.next == nil {
		return Node[T]{}.Value, false
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
