package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	if l.Head == nil {
		l.Head, l.Tail = &NodeL{Data: data}, l.Head
	} else {
		addr := &NodeL{Data: data}
		addr.Next = l.Head
		l.Head = addr
	}
}
