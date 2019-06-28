package piscine

func listReverse(l *list) {

	if l == nil {
		return
	}

	current := l.head
	prev := l.head
	prev = nil

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
}