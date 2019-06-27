package piscine

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	addr := &NodeL{Data: data}

	if l.Head == nil {
		l.Head, l.Tail = addr, l.Head
	}
	addr.Next = l.Head
	l.Head = addr
}
