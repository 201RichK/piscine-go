package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	l = ListSort(l)
	n := &NodeI{}
	n.Data = data_ref
	n.Next = nil

	if l == nil || l.Data >= n.Data {
		n.Next = l
		return n
	} else {
		temp := l
		for temp.Next != nil && temp.Next.Data < n.Data {
			temp = temp.Next
		}
		n.Next = temp.Next
		temp.Next = n
	}
	return l
}
