package piscine

func ListReverse(l *List) {

	current := l.Head
	var prev *NodeL

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return l.Head = prev
}