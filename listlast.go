package piscine

type Node struct {
	Data interface{}
	Next *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	} else {
		if l.Head.Next == nil {
			return l.Head.Data
		}
		l.Head = l.Head.Next
	}
	return l.Head.Data
}
