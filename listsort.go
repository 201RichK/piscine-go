package piscine

func listSort(l *node) *node {

	ptr := l
	if ptr == nil {
		return nil
	}
	ptr.next = listSort(ptr.next)

	if ptr.next != nil && ptr.data > ptr.next.data {
		ptr = move(ptr)
	}
	return ptr
}
