package piscine

func ListAt(l *NodeL, nbr int) *NodeL {

	index := 0

	for l != nil {
		index++
		l = l.Next
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
