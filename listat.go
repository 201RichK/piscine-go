package piscine

func ListAt(l *NodeL, nbr int) *NodeL {

	head := l
	for index := 0; head != nil; index++ {
		head = head.Next
	}

	if nbr <= index {
		for i := 0; l != nil; i++ {
			if i == nbr {
				return l
			}
			l = l.Next
		}
	}
	return nil
}
