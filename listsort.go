package piscine

func listSort(l *Node) *Node {

	ptr := l
	if ptr == nil {
		return nil
	}
	ptr.Next = listSort(ptr.Next)

	if ptr.Next != nil && ptr.data > ptr.Next.data {
		ptr = move(ptr)
	}
	return ptr
}

func move(l *Node) *Node {
	p := l
	n := l.Next
	ret := n

	for n != nil && l.Data > n.Data {
		p = n
		n = n.Next
	}
	p.Next = l
	l.Next = n
	return ret
}
