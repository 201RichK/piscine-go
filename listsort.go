package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {

	pointer := l
	if pointer == nil {
		return nil
	}
	pointer.Next = ListSort(pointer.Next)

	if pointer.Next != nil && pointer.Data > pointer.Next.Data {
		pointer = move(pointer)
	}
	return pointer
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
