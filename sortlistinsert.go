package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	n := &NodeI{Data: data_ref}
	if l == nil {
		return n
	}
	un := l
	for l.Next != nil && l.Next.Data < n.Data {
		l = l.Next
	}
	n.Next = l.Next
	l.Next = n

	return ListSort(un)
}
