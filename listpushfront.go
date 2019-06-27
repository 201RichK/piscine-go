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
	addr := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = &NodeL{Data: data}
		l.Tail = l.Head
	}

	addr.Next = l.Head
	l.Head = addr

}
