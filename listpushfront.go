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
		l.Head = &NodeL{Data: data}
		l.Tail = l.Head
	}

	&NodeL{Data: data}.Next = l.Head
	l.Head = &NodeL{Data: data}

}
