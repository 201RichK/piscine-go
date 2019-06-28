package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {

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

func move(l *NodeI) *NodeI {
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
