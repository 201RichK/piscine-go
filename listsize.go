package piscine

func ListSize(l *List) int {
	i := 0
	currentNode := l.Head

	if currentNode == nil {
		return i
	}

	for currentNode.Next != nil {
		i = i + 1
		currentNode = currentNode.Next
	}
	return i

}
