package piscine

func ListAt(l *NodeL, nbr int) *NodeL {

	head := l
   index := 0

	for head != nil;  {
    index++
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
