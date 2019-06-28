package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {

	pointer := l
	if pointer == nil {
		return nil
	}
	pointer.Next = ListSort(pointer.Next)

	if pointer.Next != nil && pointer.Data > pointer.Next.Data {
		pointer += 1
	}
	return pointer
}
