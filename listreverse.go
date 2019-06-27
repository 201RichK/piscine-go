package piscine

func ListReverse(l *List) {
	pCur := l

	var pTop *List = nil

	for {
		if pCur == nil {
			break
		}

		pTemp := pCur.Next
		pCur.Next = pTop
		pTop = pCur
		pCur = pTemp

	}
	return pTop

}
