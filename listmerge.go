package piscine

func ListMerge(l1 *List, l2 *List) {
	iterator := l1.Head
	iterator2 := l2.Head

	for iterator != nil {
		if iterator.Next == nil {
			iterator.Next = iterator2
			return
		}
		iterator = iterator.Next
	}

}
