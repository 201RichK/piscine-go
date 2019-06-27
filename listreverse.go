package piscine

//reverses the list
func ListReverse(l *List) {
	if l == nil {
		return l
	}

	p := l.Head
	if p.Next == nil {
		return p
	} else {
		newHead := ListReverse(p.Next)
		p.Next.Next = p
		p.Next = nil
		return newHead
	}
}
