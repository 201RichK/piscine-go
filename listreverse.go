package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

func ListSize(l *List) int {
	currentNode := l.Head
	i := 1

	if currentNode == nil {
		return 0
	}

	for currentNode.Next != nil {
		currentNode = currentNode.Next
		i++
	}
	return i

}
