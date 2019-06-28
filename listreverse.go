package piscine

func listReverse(l *List) {

	if l == nil {
		return
	}

	current := l.Head
	prev := l.Head
	prev = nil

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}